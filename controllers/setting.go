package controllers

import (
<<<<<<< HEAD
	"encoding/json"
	"fmt"
	"io/ioutil"
	"look/models"
	"strconv"

	"github.com/astaxie/beego"
)
=======
	"io/ioutil"
	"look/models"
	"fmt"
	"strconv"

	"github.com/astaxie/beego"
) 
>>>>>>> ce945451a69368a9de219281cf30b3fc796510bd

type SettingController struct {
	beego.Controller
}

<<<<<<< HEAD
func (this *SettingController) Get() {
	result := map[string]interface{}{
		"time":  models.HeartBeatsTime,
		"theme": nowTheme,
		"email": models.ToAddress}

	this.Data["json"] = result
	this.ServeJSON()

}

func (this *SettingController) Post() {

	var params map[string]interface{}
	result := make(map[string]map[string]interface{}, 3)

	defer this.Ctx.Request.Body.Close()
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("出了错：", err)
		}
		this.Data["json"] = result
		this.ServeJSON()
	}()
	body, _ := ioutil.ReadAll(this.Ctx.Request.Body)

	json.Unmarshal(body, &params)
	fmt.Println(string(body))
	fmt.Println(params)

	time := params["time"]
	theme := params["theme"]
	email := params["email"]

	if time != nil {
		hbtime, err := strconv.ParseInt(time.(string), 10, 64)
		if err != nil {
			fmt.Println(err)
			result["time"] = createResultMap(false, "heartbeats time should is a number")
		} else {
			models.HeartBeatsTime = hbtime
			result["time"] = createResultMap(true, "")
		}
	}

	if theme != nil {
		themeNo, err := strconv.Atoi(theme.(string))
		if err != nil || themeNo > len(themes)-1 || themeNo < 0 {
			fmt.Println(err)
			result["theme"] = createResultMap(false, "theme should is a number")
		} else {
			nowTheme = themeNo
			result["theme"] = createResultMap(true, "")
		}
	}

	if email != nil {
		to, bool := email.(string)
		if !bool || !checkEmail(to) {
			result["email"] = createResultMap(false, "mail format is error")
		} else {
			models.ToAddress = to
			result["email"] = createResultMap(true, "")
		}
	}
}

func createResultMap(s bool, m string) map[string]interface{} {
	result := make(map[string]interface{})
	result["success"] = s
	result["message"] = m
	return result

}

func checkEmail(email string) bool {
	return true
}
=======
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
>>>>>>> ce945451a69368a9de219281cf30b3fc796510bd
