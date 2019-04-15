package main

import (
	_ "firstbeego/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	orm.RegisterDataBase("default", "mysql", "root:bjtungirc@tcp(127.0.0.1:3306)/firstbeego")
	beego.Run()
}
