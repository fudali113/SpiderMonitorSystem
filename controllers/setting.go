package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"look/models"
	"strconv"

	"github.com/astaxie/beego"
)

type SettingController struct {
	beego.Controller
}
type DefalutController struct {
	beego.Controller
}

const (
	defaultT = 0
)

func (this *DefalutController) Post() {
	models.HeartbeatTime = models.DefaultHT
	models.PcDownSendEmailTime = models.DefaultPDSET
	nowTheme = defaultT
	models.ToAddress = models.ToAddress
	result := map[string]interface{}{
		"time":     strconv.FormatInt(models.HeartbeatTime, 10),
		"theme":    strconv.Itoa(nowTheme),
		"email":    models.ToAddress,
		"sendtime": strconv.FormatInt(models.PcDownSendEmailTime, 10)}

	this.Data["json"] = result
	this.ServeJSON()
}

func (this *SettingController) Get() {
	result := map[string]interface{}{
		"time":     strconv.FormatInt(models.HeartbeatTime, 10),
		"theme":    strconv.Itoa(nowTheme),
		"email":    models.ToAddress,
		"sendtime": strconv.FormatInt(models.PcDownSendEmailTime, 10)}

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
		this.Get()
	}()

	body, _ := ioutil.ReadAll(this.Ctx.Request.Body)

	json.Unmarshal(body, &params)
	fmt.Println(string(body))
	fmt.Println(params)

	time := params["time"]
	sendtime := params["sendtime"]
	theme := params["theme"]
	email := params["email"]

	if time != nil {
		hbtime, err := strconv.ParseInt(time.(string), 10, 64)
		if err != nil {
			fmt.Println(err)
			result["time"] = createResultMap(false, "heartbeats time should is a number")
		} else {
			models.HeartbeatTime = hbtime
			result["time"] = createResultMap(true, strconv.FormatInt(models.HeartbeatTime, 10))
		}
	}

	if sendtime != nil {
		pdset, err := strconv.ParseInt(sendtime.(string), 10, 64)
		if err != nil {
			fmt.Println(err)
			result["sendtime"] = createResultMap(false, "send time should is a number")
		} else {
			models.PcDownSendEmailTime = pdset
			result["sendtime"] = createResultMap(true, strconv.FormatInt(models.PcDownSendEmailTime, 10))
		}
	}

	if theme != nil {
		themeNo, err := strconv.Atoi(theme.(string))
		if err != nil || themeNo > len(themes)-1 || themeNo < 0 {
			fmt.Println(err)
			result["theme"] = createResultMap(false, "theme should is a number")
		} else {
			nowTheme = themeNo
			result["theme"] = createResultMap(true, strconv.Itoa(nowTheme))
		}
	}

	if email != nil {
		to, bool := email.(string)
		if !bool || !checkEmail(to) {
			result["email"] = createResultMap(false, "mail format is error")
		} else {
			models.ToAddress = to
			result["email"] = createResultMap(true, models.ToAddress)
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
