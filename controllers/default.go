package controllers

import (
	"io/ioutil"
	"look/models"

	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

type IndexController struct {
	beego.Controller
}

func (this *IndexController) Get() {
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
