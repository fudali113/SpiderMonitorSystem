package controllers

import (
	"look/models"

	"github.com/astaxie/beego"
)

type SSController struct {
	beego.Controller
}

func (this *SSController) Get() {
	this.Data["json"] = map[string]interface{}{
		"pschan": len(models.PS),
		"mchan":  len(models.Messages)}
	this.ServeJSON()
}
