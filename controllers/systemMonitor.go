package controllers

import (
	"look/models"

	"github.com/astaxie/beego"
)

type SMController struct {
	beego.Controller
}

func (this *SMController) Get() {
	pcid := this.Ctx.Input.Param(":pcid")
	who := this.Ctx.Input.Param(":who")
	this.Ctx.Output.Header("Content-Type", "application/json; charset=utf-8")
	body,_ := models.GetSysInfo(pcid, who)
	this.Ctx.Output.Body(body)
}

type PidInfoController struct {
	beego.Controller
}

func (this *PidInfoController) Get()  {
	pcid := this.Ctx.Input.Param(":pcid")
	pid := this.Ctx.Input.Param(":pid")
	this.Ctx.Output.Header("Content-Type", "application/json; charset=utf-8")
	body,_ := models.GetPidInfo(pcid, pid)
	this.Ctx.Output.Body(body)
}
