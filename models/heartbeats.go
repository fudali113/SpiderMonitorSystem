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

var History map[string]int64

const (
	checkTimer = time.Second * 5
)

var (
	HeartBeatsTime int64 = 5000
)

func RecordPcLastTime(pcstatus []byte) {
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
			}
			t1.Reset(time.Millisecond * time.Duration(HeartBeatsTime))

		case <-t2.C:
			t2.Reset(time.Second * 10)
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
	go checkHB()
}
