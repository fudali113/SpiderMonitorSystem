package models

import (
	"encoding/json"
	"fmt"
	"time"
)

type HeartBeats struct {
	Cid        string `json:"pc_id"`
	HeartBeats int    `json:"hb"`
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
	s := &HeartBeats{HeartBeats: -1}
	err := json.Unmarshal(pcstatus, s)
	if err != nil {
		fmt.Println(err)
	}
	if s.HeartBeats == 0 {
		fmt.Println(s.Cid)
		sendPcDown(s)
	}
	pcid := s.Cid
	HistoryData[pcid] = string(pcstatus)
	if notIn(pcid) {
		go func() {
			time.Sleep(time.Millisecond * 500)
			Messages <- string(pcstatus)
		}()
	}
	if pcid != "" {
		nowTime := time.Now().Unix()
		History[pcid] = nowTime
		sendPcDown(&HeartBeats{Cid: pcid, HeartBeats: 1})
	}
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
				if missTime > HeartBeatsTime/1000 {
					sendPcDown(&HeartBeats{Cid: k, HeartBeats: 0})
				}
				if missTime > PcDownSendEmailTime*60 {
					m := map[string]interface{}{
						"before":   missTime,
						"pc_id":    k,
						"downTime": v,
						"lastData": HistoryData[k]}

					body, _ := GetHtmlWithTpl("views/email.tpl", m)
					email := Email{To: ToAddress,
						Subject:  "haved a computer is down",
						Body:     body,
						MailType: "html"}
					SendEmail(email)
					delete(History, k)
					fmt.Println("send one email to " + ToAddress)
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
