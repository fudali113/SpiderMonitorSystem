package routers

import (
	"look/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/message", &controllers.MainController{})
	beego.Router("/ws", &controllers.WSController{})
	beego.Router("/", &controllers.IndexController{})
}
