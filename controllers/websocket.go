package controllers

import (
	"fmt"
	"look/models"

	"github.com/astaxie/beego"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type WsController struct {
	beego.Controller
}

func (this *WsController) Get() {
	ws, err := upgrader.Upgrade(this.Ctx.ResponseWriter, this.Ctx.Request, nil)
	if err != nil {
		fmt.Println(err)
	} else {
		models.Wss = append(models.Wss, ws)
	}
}
