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
	t := &mysql.TimeFilter{Start: this.GetString("start"), End: this.GetString("end")}
	fmt.Println(t)
	this.Data["json"] = mysql.GetExecAllRatio(t)
	fmt.Println(this.Data["json"])
	this.ServeJSON()
}

func (this *DataController) StepFinishRatio() {
	t := &mysql.TimeFilter{Start: this.GetString("start"), End: this.GetString("end")}
	fmt.Println(t)
	this.Data["json"] = mysql.GetStepFinish(t)
	fmt.Println(this.Data["json"])
	this.ServeJSON()
}

func (this *DataController) PcDownRatio() {
	t := &mysql.TimeFilter{Start: this.GetString("start"), End: this.GetString("end")}
	fmt.Println(t)
	this.Data["json"] = mysql.GetPcDownRatio(t)
	fmt.Println(this.Data["json"])
	this.ServeJSON()
}
