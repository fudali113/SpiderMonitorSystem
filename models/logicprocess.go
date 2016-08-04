package models

import (
	"encoding/json"
	"fmt"
	"look/mysql"
	"strconv"
	"time"

	"github.com/astaxie/beego"
)

const (
	DefaultHT    int64 = 5
	DefaultPDSET int64 = 5
	checkTimer         = time.Second * 5
)

var (
	History             = &SI64Map{M: make(map[string]int64)}
	pcipmap             = &S2Map{M: make(map[string]string)}    //pcid与ip映射
	HistoryData         = &S2Map{M: make(map[string]string)}    //pcid最后有效数据
	PS                  = make(chan []byte, 10000)              //PcStatus
	HeartbeatTime       = DefaultHT                             //检查心跳时间
	PcDownSendEmailTime = DefaultPDSET                          // 多久没有心跳发送邮件时间
	STC                 = &SSIMap{M: make(map[string]StepInfo)} //StepTimeConsuming
	MTC                 = &TrafficCount{}                       //每分钟流量计数器
)

func recordInfo(pcstatus []byte) {
	s := &PcStatus{}
	err := json.Unmarshal(pcstatus, s)
	if err != nil {
		beego.Error(err)
		return
	}

	pcid := s.PCid
	if pcid == "" {
		return
	}
	ip := s.Ip
	ss := s.SpiderStatus
	pc_exception := s.Exception
	bid := s.Bank
	nowTime := time.Now().Unix()
	nowTimeStr := time.Unix(nowTime, 0).Format("2006-01-02 15:04:05")

	History.Put(pcid, time.Now().Unix())
	pcipmap.Put(pcid, ip)

	if pc_exception != "" {
		beego.Warning("pc exception is ", pc_exception)
		go SendEmailWithMap(map[string]interface{}{
			"pcid": pcid,
			"ip":   ip,
			"pce":  pc_exception,
			"ss":   ss,
			"time": nowTimeStr,
			"data": string(pcstatus)}, "haved a pc in exception", "views/exception.tpl")
	}
	if ss == nil { //如果ss为空，返回
		return
	}
	beego.Notice(string(pcstatus))

	HistoryData.Put(pcid, string(pcstatus))
	step := -1
	if ss["step"] != nil {
		step = int(ss["step"].(float64))
	}
	sid := ""
	if ss["sid"] != nil {
		sid = ss["sid"].(string)
	}

	MTC.Incr()
	//检测爬虫每步用时
	var timeConsuming int64 = 0
	if sid != "" && step >= 0 {
		v, ok := STC.GetAndExist(sid)
		if ok {
			s := v.Step
			if s != step {
				timeConsuming = nowTime - v.Time
			} else {
				beego.Notice("return ")
				return
			}
		}
		if step == 6 {
			STC.Delete(sid) //完成第6步，删除此sid
		} else {
			STC.Put(sid, StepInfo{Time: nowTime, Step: step, Bank: bid, Pcid: pcid})
		}
	}
	if timeConsuming == 0 {
		sendMessage(pcstatus)
	} else {
		s.SpiderStatus["stc"] = timeConsuming
		jsonStr, err := json.Marshal(s)
		if err != nil {
			sendMessage(pcstatus)
		} else {
			sendMessage(jsonStr)
		}
	}
	//检测爬虫每步用时

	exception := ""
	if ss["exception"] != nil {
		exception = ss["exception"].(string)
	}

	adid := mysql.InsertAll(&mysql.All{ // 插入数据并获取插入数据id供exception插入使用
		Pcid:      pcid,
		Ip:        ip,
		Step:      step,
		Bid:       bid,
		Sid:       sid,
		Stc:       timeConsuming,
		Exception: exception,
		All:       string(pcstatus)})

	if exception != "" {
		beego.Warning("bank_status.exception:", exception)
		go func() {
			SendEmailWithMap(map[string]interface{}{
				"pcid": pcid,
				"ip":   ip,
				"pce":  pc_exception,
				"ss":   ss,
				"time": nowTimeStr,
				"data": string(pcstatus)}, "haved a spider in exception", "views/exception.tpl")
			mysql.InsertExecption(&mysql.Exception{
				Adid:      adid,
				Exception: exception})
		}()

	}
	go func() {
		mysql.IOUFinish(&mysql.Finish{
			Pcid: pcid,
			Bid:  bid,
			Sid:  sid,
			Step: step})
	}()

}

func checkHB() {
	hbs := make([]*Heartbeat, 0)
	for k, v := range History.M {
		ip := pcipmap.Get(k)
		nowTime := time.Now().Unix()
		missTime := nowTime - v
		downTimeStr := time.Unix(v, 0).Format("2006-01-02 15:04:05")
		if missTime < HeartbeatTime {
			hbs = append(hbs, &Heartbeat{PCid: k, Ip: ip, Hb: 1})
		} else if missTime >= HeartbeatTime && PcDownSendEmailTime*60 > missTime {
			hbs = append(hbs, &Heartbeat{PCid: k, Ip: ip, Hb: 0})
		} else {
			History.Delete(k)
			hbs = append(hbs, &Heartbeat{PCid: k, Ip: ip, Hb: -1})
			go mysql.InsertHB(&mysql.HB{Pcid: k, Ip: ip, Deadtime: time.Unix(v, 0)})
			go SendEmailWithMap(map[string]interface{}{
				"before":   missTime,
				"pc_id":    k,
				"downTime": downTimeStr,
				"lastData": HistoryData.Get(k)}, "haved a computer is down", "views/email.tpl")
			beego.Error("a computer is down and send email")
		}
	}
	if len(hbs) != 0 {
		hbjson, err := json.Marshal(hbs)
		if err != nil {
			beego.Error(err)
		} else {
			sendMessage(hbjson)
		}
	}
}

func trafficMonitor() { //流量统计
	mtc := MTC.GetAndReset(time.Now().Unix())
	if mtc != 0 {
		mysql.InsertTraffic(&mysql.Traffic{Count: mtc})
	}
}

func checkSpider() { //检查爬虫
	t := time.Now().Unix()
	for k, v := range STC.M {
		beego.Notice("bianli spider", k, v)
		if t-v.Time > 15 {
			v, err := json.Marshal(map[string]string{"pc_id": v.Pcid, "delete": v.Bank})
			if err != nil {
				beego.Error("删除失效Spider失败", err)
			} else {
				STC.Delete(k)
				sendMessage(v)
			}
		}
	}
}

func sendComputerStatus() {
	for k, _ := range History.M {
		body, err := GetSysInfo(k, "all")
		if err != nil {
			continue
		}
		go func() {
			perfix := []byte(fmt.Sprintf(`{"pc_id":"%s","sys":`, k))
			last := []byte(`}`)
			changeBody := append(perfix, body...)
			changeBody = append(changeBody, last...)
			sendMessage(changeBody)
		}()
		go func() {
			var data CompSysStatus
			var cpu int
			json.Unmarshal(body, &data)
			if data.Cpu != nil || len(data.Cpu) > 0 {
				cpu = data.Cpu[0]
			}
			mem, _ := strconv.Atoi(data.Mem["usedpercent"])
			mysql.InsertCS(&mysql.CompStatus{
				Pcid: k,
				Cpu:  cpu,
				Mem:  mem,
			})
		}()
	}
}

func sendMessage(hb []byte) { //发送消息
	if len(Wss) > 0 {
		select {
		case Messages <- hb:
		default:
			beego.Warning("websocket send message error : ", string(hb))
		}
	}
}

func init() {
	beego.Notice("checkHB is init")
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
	t2 := time.NewTimer(time.Second * 60)

	for {
		select {

		case <-t1.C:
			checkHB() //不能启用协程，不然可能到时map不安全与发送多天信息
			t1.Reset(time.Second * time.Duration(HeartbeatTime))

		case <-t2.C:
			trafficMonitor()
			checkSpider()
			sendComputerStatus()
			t2.Reset(time.Second * 60)
		}
	}
}
