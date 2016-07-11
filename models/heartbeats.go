package models

import (
	"encoding/json"
	"fmt"
	"look/mysql"
	"time"
)

type HeartBeats struct {
	Cid string `json:"pc_id"`
	Hb  int    `json:"hb"`
}

type PcStatus struct {
	Cid          string                 `json:"pc_id"`
	Hb           int                    `json:"hb"`
	Ip           string                 `json:"ip"`
	SpiderStatus map[string]interface{} `json:"bank_status"`
}

type StepLastInfo struct {
	Time int64
	Data string
}

var History map[string]int64
var HistoryData map[string]string
var HistoryStep map[string]StepLastInfo

const (
	checkTimer = time.Second * 5
)

var (
	HeartBeatsTime      int64 = 5000
	PcDownSendEmailTime int64 = 5
)

func RecordPcLastTime(pcstatus []byte) { //记录个pc_id发来的最后消息的时间
	s := &PcStatus{Hb: -1}
	err := json.Unmarshal(pcstatus, s)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(s)

	ss := s.SpiderStatus
	fmt.Println(ss)
	fmt.Println(ss["sid"])

	pcid := s.Cid
	ip := s.Ip
	execption := ss["execption"].(string)
	step := int(ss["step"].(float64))
	bid := ss["bid"].(string)
	sid := ss["sid"].(string)
	nowTime := time.Unix(time.Now().Unix(), 0).Format("2006-01-02 15:04:05")

	if execption != "" {
		m := map[string]interface{}{
			"pcid": pcid,
			"ip":   ip,
			"ss":   ss,
			"time": nowTime,
			"data": string(pcstatus)}

		body, _ := GetHtmlWithTpl("views/execption.tpl", m)
		email := Email{To: ToAddress,
			Subject:  "haved a spider in execption",
			Body:     body,
			MailType: "html"}

		go func() {
			SendEmail(email)
			mysql.InsertExecption(&mysql.Execption{
				Pcid:      pcid,
				Ip:        ip,
				Step:      step,
				Bid:       bid,
				Execption: execption,
				Data:      string(pcstatus)})
		}()

	} else {
		go mysql.InsertAll(&mysql.All{
			Pcid: pcid,
			Ip:   ip,
			Step: step,
			Bid:  bid,
			Sid:  sid,
			All:  string(pcstatus)})
	}

	if pcid == "" {
		return
	}

	if s.Hb == 0 {
		fmt.Println(s.Cid)
		sendPcDown(&HeartBeats{Cid: pcid, Hb: 0})
	}

	if notIn(pcid) {
		go func() {
			time.Sleep(time.Millisecond * 500)
			Messages <- string(pcstatus)
		}()
	}

	History[pcid] = time.Now().Unix()
	HistoryData[pcid] = string(pcstatus)
	sendPcDown(&HeartBeats{Cid: pcid, Hb: 1})
}

func notIn(id string) bool {
	var count = 0
	for k, _ := range History {
		if k == id {
			count = count + 1
		}
	}
	return count == 0
}

func checkHB() {
	t1 := time.NewTimer(checkTimer)
	t2 := time.NewTimer(time.Second * 10)

	for {
		select {

		case <-t1.C:
			for k, v := range History {
				nowTime := time.Now().Unix()
				missTime := nowTime - v
				downTimeStr := time.Unix(v, 0).Format("2006-01-02 15:04:05")

				if missTime > HeartBeatsTime/1000 {
					sendPcDown(&HeartBeats{Cid: k, Hb: 0})
				}
				if missTime > PcDownSendEmailTime*60 {
					go func() {
						m := map[string]interface{}{
							"before":   missTime,
							"pc_id":    k,
							"downTime": downTimeStr,
							"lastData": HistoryData[k]}

						body, _ := GetHtmlWithTpl("views/email.tpl", m)
						email := Email{To: ToAddress,
							Subject:  "haved a computer is down",
							Body:     body,
							MailType: "html"}
						SendEmail(email)
						delete(History, k)
					}()

				}
			}
			t1.Reset(time.Millisecond * time.Duration(HeartBeatsTime))

		case <-t2.C:
			t2.Reset(time.Minute * time.Duration(PcDownSendEmailTime))
		}
	}
}

func sendPcDown(hb *HeartBeats) {
	hbjson, _ := json.Marshal(hb)
	Messages <- string(hbjson)
}

func init() {
	fmt.Println("checkHB is init")
	History = make(map[string]int64)
	HistoryData = make(map[string]string)
	go checkHB()
}
