package routers

import (
	"hellogo/controllers"
	"github.com/astaxie/beego"
	"fmt"
)

func init() {
	beego.Router("/", &controllers.MainController{})
}

type UnRoute struct {

}

func unDelete() {
	fmt.Print(".....")
}