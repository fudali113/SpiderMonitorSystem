package controllers

import (
	"fmt"
	"io/ioutil"
	"look/models"
	"strings"

	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

type IndexController struct {
	beego.Controller
}

func (this *IndexController) Get() {
	addrs := strings.Split(this.Ctx.Request.RemoteAddr, "::1")
	port := beego.AppConfig.String("httpport")
	wsip := models.LocalIp
	if len(addrs) > 1 {
		wsip = "localhost"
	}
	this.Data["ip"] = wsip + ":" + port
	this.TplName = "index.tpl"
}

func (this *MainController) Post() {
	defer this.Ctx.Request.Body.Close()
	body, err := ioutil.ReadAll(this.Ctx.Request.Body)
	fmt.Println(len(body))
	fmt.Println(string(body))

	if err == nil && body != nil {
		select {
		case models.PS <- body:
			this.Data["json"] = "ok"
		default:
			this.Data["json"] = "server is busy"
		}
	} else {
		this.Data["json"] = "fail"
	}
	this.ServeJSON()
}
