package routers

import (
	"mygolang/mydemo/mydemo/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})


    beego.Router("/register", &controllers.UsersController{}, "*:Register")
    beego.Router("/login", &controllers.UsersController{}, "*:Login")
    beego.Router("/update", &controllers.UsersController{}, "*:Update")
    beego.Router("/upload", &controllers.UsersController{}, "*:Upload")

    beego.Router("/register/register", &controllers.UsersController{}, "Post:PostRegister")
    beego.Router("/login/login", &controllers.UsersController{}, "Post:PostLogin")
    beego.Router("/update/update", &controllers.UsersController{}, "Post:PostUpdate")
    beego.Router("/upload/upload", &controllers.UsersController{}, "Post:PostUpload")

    beego.Router("/ws/join", &controllers.WebSocketController{}, "*:Join")
    beego.Router("/ws", &controllers.WebSocketController{}, "*:Post")
}
