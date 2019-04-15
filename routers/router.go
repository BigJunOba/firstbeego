package routers

import (
	"firstbeego/controllers"
	"github.com/astaxie/beego"
)

func init() {
	//beego.Router("/", &controllers.MainController{})
	beego.Include(&controllers.UserController{})
}
