package models

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"

	"github.com/astaxie/beego"
)

var (
	port = "8888"
)

func GetSysInfo(pcid, who string) []byte {
	pcip := GetPcIP(pcid)
	if pcip == "" {
		return []byte(`{"err":"error pcid"}`)
	}
	return GetResponse(CreatUrl(pcip, port, who))
}

func GetResponse(url string) []byte {
	return []byte(fmt.Sprintf(`{"cpu":[%d],"mem":{"usedpercent":%d}}`, rand.Intn(30), 40+rand.Intn(10)))
	res, err := http.Get(url)
	if err != nil {
		beego.Error("获取数据   info/all  get请求", err)
		return nil
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		beego.Error(" 获取数据  info/all  response body 读取", err)
		return nil
	}
	var oo map[string]interface{}
	json.Unmarshal(body, &oo)
	fmt.Println(oo)
	return body
}

func CreatUrl(host, port, who string) string {
	return fmt.Sprintf("http://%s:%s/info/%s", host, port, who)
}

func GetPcIP(pcid string) string {
	return "192.168.0.113"
	//return pcipmap.Get(pcid)
}
