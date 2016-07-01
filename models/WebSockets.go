package models

import (
	"fmt"

	"github.com/gorilla/websocket"
)

var Wss = make([]*websocket.Conn, 5)
var Messages = make(chan string)

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
						panic(err)
					}
				}
			}
		}
	}
}

func init() {
	go myinit()
}
