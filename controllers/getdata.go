package controllers

import (
	"fmt"

	"look/mysql"

	"github.com/astaxie/beego"
)

type DataController struct {
	beego.Controller
}

func (this *DataController) StepExecAllRatio() {
	this.Data["json"] = mysql.GetExecAllRatio()
	fmt.Println(this.Data["json"])
	this.ServeJSON()
}

func (this *DataController) StepFinishRatio() {
	this.Data["json"] = mysql.GetStepFinish()
	fmt.Println(this.Data["json"])
	this.ServeJSON()
}
