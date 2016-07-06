package controllers

import (
	"io/ioutil"
	"look/models"
	"strings"
	"fmt"

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
	go models.RecordPcLastTime(body)
	content := string(body)

	fmt.Println(content)

	if err == nil && content != "" {
		models.Messages <- content
		this.Data["json"] = "ok"
	} else {
		this.Data["json"] = "fail"
	}

	this.ServeJSON()
}
