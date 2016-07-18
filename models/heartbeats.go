package models

import (
	"encoding/json"
	"fmt"
	"look/mysql"
	"time"
)

type Heartbeat struct {
	PCid string `json:"pc_id"`
	Hb   int    `json:"hb"`
}

type PcStatus struct {
	PCid         string                 `json:"pc_id"`
	Ip           string                 `json:"pc_ip"`
	Bank         string                 `json:"bank_name"`
	Execption    string                 `json:"execption"`
	SpiderStatus map[string]interface{} `json:"bank_status"`
}

const (
	DefaultHT    int64 = 5
	DefaultPDSET int64 = 5
)

var (
	History             = make(map[string]int64)
	HistoryData         = make(map[string]string)
	PS                  = make(chan []byte, 10000)
	HeartbeatTime       = DefaultHT
	PcDownSendEmailTime = DefaultPDSET
)

const (
	checkTimer = time.Second * 5
)

func recordInfo(pcstatus []byte) { //记录个pc_id发来的最后消息的时间
	s := &PcStatus{}
	err := json.Unmarshal(pcstatus, s)
	if err != nil {
		fmt.Println(err)
		return
	}

	pcid := s.PCid
	if pcid == "" {
		return
	}
	ip := s.Ip
	ss := s.SpiderStatus
	pc_execption := s.Execption
	bid := s.Bank
	nowTime := time.Unix(time.Now().Unix(), 0).Format("2006-01-02 15:04:05")

	History[pcid] = time.Now().Unix()

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

	HistoryData[pcid] = string(pcstatus)
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
	if len(Wss) > 0 {
		select {
		case Messages <- pcstatus:
			//fmt.Print("websocket get spider status : ")
		default:
			fmt.Println("websocket send ss error ")
		}
	}

	execption := ""
	if ss["execption"] != nil {
		execption = ss["execption"].(string)
	}

	if execption != "" {
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

//func notIn(id string) bool {
//	var count = 0
//	for k, _ := range History {
//		if k == id {
//			count = count + 1
//		}
//	}
//	return count == 0
//}

func checkHB() {
	for k, v := range History {
		nowTime := time.Now().Unix()
		missTime := nowTime - v
		downTimeStr := time.Unix(v, 0).Format("2006-01-02 15:04:05")

		if missTime < HeartbeatTime {
			sendHbMessage(&Heartbeat{PCid: k, Hb: 1})
		} else if missTime >= HeartbeatTime && PcDownSendEmailTime*60 > missTime {
			sendHbMessage(&Heartbeat{PCid: k, Hb: 0})
		} else {
			delete(History, k)
			go mysql.InsertHB(&mysql.HB{Pcid: k, Deadtime: time.Unix(v, 0)})
			go SendEmailWithMap(map[string]interface{}{
				"before":   missTime,
				"pc_id":    k,
				"downTime": downTimeStr,
				"lastData": HistoryData[k]}, "haved a computer is down", "views/email.tpl")
		}
	}
}

func sendHbMessage(hb *Heartbeat) {
	hbjson, _ := json.Marshal(hb)
	select {
	case Messages <- hbjson:
		//fmt.Print("websocket get heartbeat : ")
	default:
		fmt.Println("websocket send hb error ")
	}
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
			go recordInfo(ps)
		}
	}
}

func check() {
	t1 := time.NewTimer(checkTimer)
	t2 := time.NewTimer(time.Second * 10)

	for {
		select {

		case <-t1.C:
			checkHB() //不能启用协程，不然可能到时map不安全与发送多天信息
			t1.Reset(time.Second * time.Duration(HeartbeatTime))

		case <-t2.C:
			t2.Reset(time.Minute * time.Duration(PcDownSendEmailTime))
		}
	}
}
