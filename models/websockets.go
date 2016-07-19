package models

import (
	"fmt"
	"net"
	"strings"
	"time"

	"github.com/gorilla/websocket"
)

const ()

var (
	LocalIp  = getLocalIp()
	Wss      = make([]*websocket.Conn, 0)
	Messages = make(chan []byte, 10000)
)

func myinit() {
	fmt.Println("init models")
	for {
		select {
		case message := <-Messages:
			before := time.Now().UnixNano()
			send(0, message)
			after := time.Now().UnixNano()
			fmt.Printf("time consuming : %d ws -> %d ns \n", len(Wss), after-before)
		}
	}
}

func send(j int, m []byte) {
	for i := j; i < len(Wss); i++ {
		conn := Wss[i]
		if conn != nil {
			err := conn.WriteMessage(websocket.TextMessage, m)
			if err != nil {
				go conn.Close()
				fmt.Println("----->一个websocket退出了连接")
				fmt.Println(err)
				if i == len(Wss)-1 {
					Wss = Wss[:i] //管理websocket连接数组，清除已断开的连接
				} else {
					Wss = append(Wss[:i], Wss[i+1:]...)
					send(i, m)
				}
			}
		}
		break
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
	go myinit()
}
