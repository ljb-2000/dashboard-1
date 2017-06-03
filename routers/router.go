package routers

import (
	"dashboard/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{}, "*:Dashboard")
	beego.Router("/images", &controllers.MainController{}, "*:Images")
	beego.Router("/dashboard", &controllers.MainController{}, "*:Dashboard")
	beego.Router("/login", &controllers.MainController{}, "*:Login")
	beego.Router("/app_manage", &controllers.MainController{}, "*:Product")
	beego.Router("/app_manage/service", &controllers.MainController{}, "*:Index")
	beego.Router("/app_manage/container", &controllers.MainController{}, "*:Container")
	beego.Router("/app_manage/configs", &controllers.MainController{}, "*:Configs")
}