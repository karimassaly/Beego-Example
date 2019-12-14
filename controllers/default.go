package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	beego.SetStaticPath("js", "views/Main/js")
	beego.SetStaticPath("css", "views/Main/css")
	beego.SetStaticPath("fonts", "views/Main/fonts")
	beego.SetStaticPath("vendor", "views/Main/vendor")
	beego.SetStaticPath("images", "views/Main/images")

	c.TplName = "Main/index.html"

}
