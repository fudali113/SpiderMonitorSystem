package controllers

import (
	"fmt"
	"look/models"
	"time"

	"github.com/astaxie/beego"
)

type SMController struct {
	beego.Controller
}

func (this *SMController) Get() {
	pcid := this.Ctx.Input.Param(":pcid")
	who := this.Ctx.Input.Param(":who")
	this.Ctx.Output.Header("Content-Type", "application/json; charset=utf-8")
	this.Ctx.Output.Body(models.GetSysInfo(pcid, who))
}
