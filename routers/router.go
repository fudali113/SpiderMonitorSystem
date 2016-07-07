package routers

import (
	"look/controllers"

	"github.com/astaxie/beego"
)

func init() {
<<<<<<< HEAD
	beego.Router("/theme", &controllers.ThemeController{})
=======
>>>>>>> ce945451a69368a9de219281cf30b3fc796510bd
	beego.Router("/setting", &controllers.SettingController{})
	beego.Router("/message", &controllers.MainController{})
	beego.Router("/ws", &controllers.WSController{})
	beego.Router("/", &controllers.IndexController{})
}
