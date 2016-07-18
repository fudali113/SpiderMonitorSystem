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
	Wss      = make([]*websocket.Conn, 5)
	Messages = make(chan []byte, 10000)
)

func myinit() {
	fmt.Println("init models")
	for {
		select {
		case message := <-Messages:
			before := time.Now().UnixNano()
			for i, _ := range Wss {
				conn := Wss[i]
				if conn != nil {
					err := conn.WriteMessage(websocket.TextMessage, message)
					if err != nil {
						conn.Close()
						fmt.Println("----->一个websocket退出了连接")
						fmt.Println(err)
						Wss = append(Wss[:i], Wss[i+1:]...) //管理websocket连接数组，清除已断开的连接
						if len(message) > 80 {
							//如果消息不为心跳信息 重发消息 保证每个连接收到消息
							Messages <- message
						}
						break //退出循环，避免数组下标异常
					}
				}
			}
			after := time.Now().UnixNano()
			fmt.Printf("time consuming : %d ns \n", after-before)
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
	go myinit()
}
