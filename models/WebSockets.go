package models

import (
	"fmt"

	"github.com/gorilla/websocket"
)

var Wss = make([]*websocket.Conn, 50)
var Messages = make(chan string, 1000)

func myinit() {
	fmt.Println("init models")
	for {
		select {
		case message := <-Messages:
			for i, _ := range Wss {
				conn := Wss[i]
				if conn != nil {
					err := conn.WriteMessage(websocket.TextMessage, []byte(message))
					if err != nil {
						fmt.Println(err)
						conn.Close()
						k := i + 1
						Wss = append(Wss[:i], Wss[k:]...) //管理websocket连接数组，清除已断开的连接
						Messages <- message               //保证消息不丢失
						break                             //退出循环，避免数组下标异常
					}
				}
			}
		}
	}
}

func init() {
	go myinit()
}
