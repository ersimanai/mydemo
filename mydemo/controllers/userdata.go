package controllers

import (
	"github.com/astaxie/beego"
	_"github.com/go-sql-driver/mysql"
	_"fmt"

	"mydemo/mydemo/models"
)

type UsersController struct {
	beego.Controller
}

func init() {

}

func (c *UsersController) PostLogin() {
	name := c.GetString("username")
	password := c.GetString("password")

	u := new(models.Userdata)
	u.Name = name 
	u.Password = password
	models.Register(u)

	c.Ctx.WriteString("登陆成功")
}

func(c *UsersController) PostRegister() {
	name := c.GetString("username")
	password := c.GetString("password")

	u := new(models.Userdata)
	u.Name = name 
	u.Password = password
	models.Register(u)

}

func(c *UsersController) PostUpdate() {

}

func (c *UsersController) PostUpload() {

}