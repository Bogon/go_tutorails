package routers

import (
	"classOne/controllers"
	"github.com/astaxie/beego"
)

func init() {

    beego.Router("/", &controllers.MainController{})

    /// register
    beego.Router("/register", &controllers.RegisterController{}, "get:ShowReg")

}
