package main

import (
	"Example/database"
	_ "Example/routers"

	"github.com/astaxie/beego"
)

func main() {
	database.InitDB()
	beego.Run()
}
