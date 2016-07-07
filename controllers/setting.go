package controllers

import (
	"io/ioutil"
	"look/models"
	"fmt"
	"strconv"

	"github.com/astaxie/beego"
) 

type SettingController struct {
	beego.Controller
}

func (this *SettingController) Post() {
	result := make(map[string]interface{})

	hbTime := this.GetString("hbtime")

	body, err := ioutil.ReadAll(this.Ctx.Request.Body)
	fmt.Println(string(body))
	
	
	hbtime,err := strconv.Atoi(hbTime)
	if err != nil {
		result["success"] = false
		result["message"] = "heartbeats time should is a number"
		this.Data["json"] = result
		this.ServeJSON()
		return
	}
	models.HeartBeatsTime = hbtime
}