package controllers

import (
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
	this.Ctx.Request.Header.Set("Access-Control-Allow-Origin","*")
	origin := this.Ctx.Request.Header.Get("Origin")
	beego.Info(origin)
	this.Ctx.Request.Header.Del("Origin")
	ws, err := upgrader.Upgrade(this.Ctx.ResponseWriter, this.Ctx.Request, nil)
	if err != nil {
		beego.Error("websocket连接出错:",err)
	} else {
		models.Wss = append(models.Wss, ws)
	}
}
