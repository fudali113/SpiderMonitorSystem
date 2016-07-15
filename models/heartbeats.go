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
	Ip           string                 `json:"pc_ip"`
	Bank         string                 `json:"bank_name"`
	Execption    string                 `json:"execption"`
	SpiderStatus map[string]interface{} `json:"bank_status"`
}

var (
	History                   = make(map[string]int64)
	HistoryData               = make(map[string]string)
	PS                        = make(chan []byte, 10000)
	HeartBeatsTime      int64 = 5
	PcDownSendEmailTime int64 = 5
)

const (
	checkTimer = time.Second * 5
)

func RecordPcLastTime(pcstatus []byte) { //记录个pc_id发来的最后消息的时间
	s := &PcStatus{Hb: -1}
	err := json.Unmarshal(pcstatus, s)
	if err != nil {
		fmt.Println(err)
		return
	}

	pcid := s.Cid
	ip := s.Ip
	ss := s.SpiderStatus
	pc_execption := s.Execption
	bid := s.Bank
	nowTime := time.Unix(time.Now().Unix(), 0).Format("2006-01-02 15:04:05")

	if pcid == "" {
		return
	}

	History[pcid] = time.Now().Unix()
	HistoryData[pcid] = string(pcstatus)

	if pc_execption != "" {
		go SendEmailWithMap(map[string]interface{}{
			"pcid": pcid,
			"ip":   ip,
			"pce":  pc_execption,
			"ss":   ss,
			"time": nowTime,
			"data": string(pcstatus)}, "haved a pc in execption", "views/execption.tpl")
	}
	if ss == nil {
		return
	}

	step := -1
	if ss["step"] != nil {
		step = int(ss["step"].(float64))
	} else {
		return
	}
	sid := ""
	if ss["sid"] != nil {
		sid = ss["sid"].(string)
	} else {
		return
	}

	select {
	case Messages <- pcstatus:
		fmt.Println("websocket 获得信息")
	default:
		fmt.Println("websocket 处理消息繁忙")
	}

	execption := ""
	if ss["execption"] != nil {
		execption = ss["execption"].(string)
	}

	if execption != "" || pc_execption != "" {
		go func() {
			SendEmailWithMap(map[string]interface{}{
				"pcid": pcid,
				"ip":   ip,
				"pce":  pc_execption,
				"ss":   ss,
				"time": nowTime,
				"data": string(pcstatus)}, "haved a spider in execption", "views/execption.tpl")
			mysql.InsertExecption(&mysql.Execption{
				Pcid:      pcid,
				Ip:        ip,
				Step:      step,
				Bid:       bid,
				Execption: execption,
				Data:      string(pcstatus)})
		}()

	}

	go func() {
		mysql.InsertAll(&mysql.All{
			Pcid:      pcid,
			Ip:        ip,
			Step:      step,
			Bid:       bid,
			Sid:       sid,
			Execption: execption,
			All:       string(pcstatus)})

		mysql.IOUFinish(&mysql.Finish{
			Pcid: pcid,
			Bid:  bid,
			Sid:  sid,
			Step: step})
	}()

	//	if notIn(pcid) {
	//		go func() {
	//			time.Sleep(time.Millisecond * 500)
	//			Messages <- pcstatus
	//		}()
	//	}
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
	for k, v := range History {
		nowTime := time.Now().Unix()
		missTime := nowTime - v
		downTimeStr := time.Unix(v, 0).Format("2006-01-02 15:04:05")

		if missTime < HeartBeatsTime {
			sendPcDown(&HeartBeats{Cid: k, Hb: 1})
		} else if missTime >= HeartBeatsTime && PcDownSendEmailTime*60 > missTime {
			sendPcDown(&HeartBeats{Cid: k, Hb: 0})
		} else {
			mysql.InsertHB(&mysql.HB{Pcid: k, Deadtime: time.Unix(v, 0)})
			SendEmailWithMap(map[string]interface{}{
				"before":   missTime,
				"pc_id":    k,
				"downTime": downTimeStr,
				"lastData": HistoryData[k]}, "haved a computer is down", "views/email.tpl")
			delete(History, k)
		}
	}
}

func sendPcDown(hb *HeartBeats) {
	hbjson, _ := json.Marshal(hb)
	Messages <- hbjson
}

func init() {
	fmt.Println("checkHB is init")
	go record()
	go check()
}

func record() {
	for {
		select {
		case ps := <-PS:
			go RecordPcLastTime(ps)
		}
	}
}

func check() {
	t1 := time.NewTimer(checkTimer)
	t2 := time.NewTimer(time.Second * 10)

	for {
		select {

		case <-t1.C:
			go checkHB()
			t1.Reset(time.Second * time.Duration(HeartBeatsTime))

		case <-t2.C:
			t2.Reset(time.Minute * time.Duration(PcDownSendEmailTime))
		}
	}
}
