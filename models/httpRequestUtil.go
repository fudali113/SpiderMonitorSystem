package models

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/astaxie/beego"
)

var (
	port = beego.AppConfig.String("performance.monitor.port")
)

func GetSysInfo(pcid, who string) ([]byte, error) {
	if host == "" {
		return []byte(`{"err":"error pcid"}`) , fmt.Errorf("pcid can`t empty")
	}
	return GetResponse(CreatUrl(pcid, port, who))
}

func GetPidInfo(pcid,pid string) ([]byte ,error) {
	host := GetPcIP(pcid)
	url := fmt.Sprintf("http://%s:%s/info/proc/port?pid=%s",host,port,pid)
	return GetResponse(url)
}

func GetResponse(url string) ([]byte, error) {
	begin := time.Now().Unix()
	res, err := http.Get(url)
	if err != nil {
		beego.Error("获取数据   info/all  get请求", err)
		return []byte(`{"err":"error get request"}`), err
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		beego.Error(" 获取数据  info/all  response body 读取", err)
		return []byte(`{"err":"error ioutil handle"}`), err
	}
	end := time.Now().Unix()
	beego.Notice(string(body))
	beego.Notice(fmt.Sprintf("%s : time comsuming -------------> %d s ", url, end-begin))
	return body, nil
}

func CreatUrl(pcid, port, who string) string {
	return fmt.Sprintf("http://%s:%s/info/%s", GetPcIP(pcid), port, who)
}

func GetPcIP(pcid string) string {
	return "192.168.0.113"
	//return pcipmap.Get(pcid)
}
