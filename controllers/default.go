package controllers

import (
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
	addrs := strings.Split(this.Ctx.Request.RemoteAddr,"::1")

	wsip := models.LocalIp
	if len(addrs) > 1 {
		wsip = "localhost"
	}
	this.Data["ip"] = wsip
	this.TplName = "index.tpl"
}

func (this *MainController) Post() {
	defer this.Ctx.Request.Body.Close()
	body, err := ioutil.ReadAll(this.Ctx.Request.Body)
	content := string(body)

	if err == nil && content != "" {
		models.Messages <- content
		this.Data["json"] = "ok"
	} else {
		this.Data["json"] = "fail"
	}

	this.ServeJSON()
}
