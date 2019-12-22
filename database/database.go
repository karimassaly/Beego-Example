package database

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq"
)

type UserForm struct {
	Id       int    `orm:"auto";primary_key`
	Name     string `json:"name" orm:"size(40)`
	Email    string `json:"email" orm:"size(64);unique"`
	Password string `json:"password"`
}

type Orm struct {
	CreatedORM orm.Ormer
}

func InitDB() {
	orm.RegisterDriver("postgres", orm.DRPostgres)
	orm.RegisterDataBase("default", "postgres", "user=doimo password=2019doimo host=127.0.0.1 port=5432 dbname=mydb sslmode=disable")
	orm.RegisterModel(new(UserForm))
	orm.RunSyncdb("default", false, true)
}
