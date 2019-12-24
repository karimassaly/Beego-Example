package controllers

import (
	"github.com/astaxie/beego"
)

type HomeController struct {
	beego.Controller
}

func (c *HomeController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	beego.SetStaticPath("js", "views/Home/js")
	beego.SetStaticPath("css", "views/Home/css")
	beego.SetStaticPath("fonts", "views/Home/fonts")
	beego.SetStaticPath("vendor", "views/Home/vendor")
	beego.SetStaticPath("images", "views/Home/images")

	c.TplName = "Home/index.html"

}
