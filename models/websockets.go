package models

import (
	"fmt"
	"net"
	"strings"

	"github.com/gorilla/websocket"
)

var LocalIp string
var Wss = make([]*websocket.Conn, 50)
var Messages = make(chan []byte, 10000)

func myinit() {
	fmt.Println("init models")
	for {
		select {
		case message := <-Messages:
			for i, _ := range Wss {
				conn := Wss[i]
				if conn != nil {
					err := conn.WriteMessage(websocket.TextMessage, message)
					if err != nil {
						fmt.Println(err)
						conn.Close()
						k := i + 1
						Wss = append(Wss[:i], Wss[k:]...) //管理websocket连接数组，清除已断开的连接
						//Messages <- message               //保证消息不丢失
						break //退出循环，避免数组下标异常
					}
				}
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
	LocalIp = getLocalIp()
	go myinit()
}
