package routers

import (
	"mydemo/mydemo/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/register", &controllers.UsersController{}, "post:PostRegister")
    beego.Router("/login", &controllers.UsersController{}, "post:PostLogin")
    beego.Router("/update", &controllers.UsersController{}, "post:PostUpdate")
    beego.Router("/upload", &controllers.UsersController{}, "post:PostUpload")

}
