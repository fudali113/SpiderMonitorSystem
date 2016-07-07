package controllers

import (
	"github.com/astaxie/beego"
)

var (
	themes = []string{
		"computer01",
		"computer02",
		"computer03",
		"computer04"}

	nowTheme = 0
)

type ThemeController struct {
	beego.Controller
}

func (this *ThemeController) Get() {
	themeAddress := "/static/html/" + themes[nowTheme] + ".html"
	this.Ctx.Redirect(302, themeAddress)
}
