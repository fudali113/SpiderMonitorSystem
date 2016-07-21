package models

import (
	"fmt"
	"net"
	"strings"
	"time"

	"github.com/astaxie/beego"
	"github.com/gorilla/websocket"
)

const (
	writeWait = 1 * time.Second
)

var (
	LocalIp  = getLocalIp()
	Wss      = make([]*websocket.Conn, 0)
	Messages = make(chan []byte, 10000)
)

func websocketMessageFS() {
	for {
		select {
		case message := <-Messages:
			before := time.Now().UnixNano()
			send(0, message)
			after := time.Now().UnixNano()
			beego.Notice("time consuming : ", len(Wss), " ws  -->  ", after-before, " ns")
		}
	}
}

func send(j int, m []byte) {
	beego.Informational("send func start : begin with ", j, " , wss length is ", len(Wss))
	for i := j; i < len(Wss); i++ {
		conn := Wss[i]
		if conn != nil {
			conn.SetWriteDeadline(time.Now().Add(writeWait))
			err := conn.WriteMessage(websocket.TextMessage, m)
			if err != nil {
				go conn.Close()
				beego.Error("a websocket conn is close , err is ", err.Error())
				if i == len(Wss)-1 {
					Wss = Wss[:i] //管理websocket连接数组，清除已断开的连接
				} else {
					Wss = append(Wss[:i], Wss[i+1:]...)
					send(i, m)
				}
				break
			}
		}
	}
}

func getLocalIp() string {
	var address string
	conn, err := net.Dial("udp", "baidu.com:80")
	defer conn.Close()
	if err != nil {
		fmt.Println(err)
		address = "localhost"
	} else {
		address = strings.Split(conn.LocalAddr().String(), ":")[0]
	}
	return address
}

func init() {
	beego.Informational("init websocketMessageFS")
	go websocketMessageFS()
}
