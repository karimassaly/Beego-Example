package controllers

import "github.com/astaxie/beego"

import "encoding/json"

type RegisterController struct {
	beego.Controller
}

type UserForm struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (this *RegisterController) Post() {
	var req UserForm
	json.Unmarshal(this.Ctx.Input.RequestBody, &req)
	println(req.Email)
}
