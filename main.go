package main

import (
	_ "look/models"
	_ "look/routers"

	"github.com/astaxie/beego"
)

func main() {
	beego.Notice("begin")
	beego.SetLogger("file", `{"filename":"logs/monitor.log"}`)
	beego.SetLevel(beego.LevelNotice)
	beego.Run()
}
