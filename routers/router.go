package routers

import (
	"look/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/theme", &controllers.ThemeController{})
	beego.Router("/setting", &controllers.SettingController{})
	beego.Router("/setting/default", &controllers.DefalutController{})
	beego.Router("/message", &controllers.MainController{})
	beego.Router("/ws", &controllers.WSController{})
	beego.Router("/", &controllers.IndexController{})
}
