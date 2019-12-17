package controllers

import (
	"Example/database"
	"encoding/json"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type RegisterController struct {
	beego.Controller
}

func (this *RegisterController) Post() {
	var req database.UserForm
	json.Unmarshal(this.Ctx.Input.RequestBody, &req)
	println(req.Email)
	println(req.Password)
	o := orm.NewOrm()
	o.Using("mydb")
	if _, err := o.Insert(&req); err != nil {
		beego.ErrorHandler("303", "Email already exists")
	}
}
