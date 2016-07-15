package main

import (
	"fmt"

	_ "look/models"
	_ "look/routers"

	"github.com/astaxie/beego"
)

func main() {
	fmt.Println("begin")
	beego.Run()
}
