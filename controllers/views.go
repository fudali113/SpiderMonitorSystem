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

func (this *MainController) Post() {
	defer this.Ctx.Request.Body.Close()
	body, err := ioutil.ReadAll(this.Ctx.Request.Body)
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

type Index2Controller struct {
	beego.Controller
}

func (this *Index2Controller) Get() {
	addrs := strings.Split(this.Ctx.Request.RemoteAddr, "::1")
	port := beego.AppConfig.String("httpport")
	wsip := models.LocalIp
	if len(addrs) > 1 {
		wsip = "localhost"
	}
	this.Data["ip"] = wsip + ":" + port
	this.TplName = "index2.tpl"
}

type PcInfoController struct {
	beego.Controller
}

func (this *PcInfoController) Get() {
	this.Data["pcid"] = this.Ctx.Input.Param(":pcid")
	this.TplName = "pcinfo.tpl"
}
