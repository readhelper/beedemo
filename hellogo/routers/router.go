package routers

import (
	"hellogo/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/rest/get", &controllers.MainController{},"get:RestGet")
	beego.Router("/rest/post", &controllers.MainController{},"get:RestPost")
	beego.Router("/rest/do", &controllers.MainController{},"get:RestDo")
	beego.Router("/rest/default", &controllers.MainController{},"get:RestDefault")
}
