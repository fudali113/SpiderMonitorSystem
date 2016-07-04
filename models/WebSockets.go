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
						Wss = append(Wss[:i], Wss[k:]...)
						Messages <- message
						break
					}
				}
			}
		}
	}
}

func init() {
	go myinit()
}
