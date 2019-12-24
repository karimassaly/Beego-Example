package controllers

import (
	"Example/database"
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type LoginController struct {
	beego.Controller
}

func (this *LoginController) Post() {
	var req database.UserForm
	var params []orm.Params
	req.Email = this.GetString("Email")
	req.Password = this.GetString("password")
	o := orm.NewOrm()
	o.Using("mydb")
	query := o.QueryTable("user_form")
	println("Success with my query", query)
	if query.Filter("email", req.Email).Exist() == false {
		this.Ctx.Output.SetStatus(404)
		this.Data["json"] = "Email doesn't exist"
		this.ServeJSON()
		return
	}

	o.Raw("SELECT * FROM user_form WHERE email = ?", req.Email).Values(&params)
	AddressPass := params[0]["password"]
	RealPass := fmt.Sprintf("%v", AddressPass)
	if RealPass != req.Password {
		this.Ctx.Output.SetStatus(309)
		this.Data["json"] = "Password does not match"
		this.ServeJSON()
		return

	} else {
		this.Ctx.Output.SetStatus(200)
		this.Data["json"] = "Login Successful"
		this.ServeJSON()
		return
	}
}
