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
	//defer this.Ctx.Request.Body.close()
	body, _ := ioutil.ReadAll(this.Ctx.Request.Body)
	models.Messages <- string(body)
}
