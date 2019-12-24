package controllers

import (
	"Example/database"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type RegisterController struct {
	beego.Controller
}

func (this *RegisterController) Post() {
	var req database.UserForm
	req.Email = this.GetString("Email")
	req.Password = this.GetString("password")
	req.Name = this.GetString("name")
	o := orm.NewOrm()
	o.Using("mydb")
	if _, err := o.Insert(&req); err != nil {
		this.Ctx.Output.SetStatus(309)
		this.Data["json"] = "Email already used"
		this.ServeJSON()
		return
	}
	this.Ctx.Output.SetStatus(200)
	this.Data["json"] = "Registration Successful"
	this.ServeJSON()
	return
}
