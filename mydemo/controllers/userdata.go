package controllers

import (
	"github.com/astaxie/beego"
	_"github.com/go-sql-driver/mysql"
	"fmt"

	"mydemo/mydemo/models"
)

const (
	Success = iota
	Failed
	Error

)


type UsersController struct {
	beego.Controller
}

func init() {

}

func (c *UsersController) Register() {
	fmt.Println("注册")
	c.TplName = "register.tpl"
}
func(c *UsersController) PostRegister() {

	name := c.GetString("username")
	password := c.GetString("password")

	u := new(models.Userdata)
	u.Name = name 
	u.Password = password
	b := models.Register(u)

	if !b {
		c.Ctx.WriteString("注册失败")
		return
	}
	c.Ctx.WriteString("注册成功")

}

func (c *UsersController) Login() {
	fmt.Println("登录")
	c.TplName = "login.tpl"

}
func (c *UsersController) PostLogin() {
	name := c.GetString("username")
	password := c.GetString("password")

	u := new(models.Userdata)
	u.Name = name 
	u.Password = password
	b := models.Login(u)

	if !b {
		c.Ctx.WriteString("用户或者密码错误")
		return
	}
	c.Ctx.WriteString("登陆成功")
}



func(c *UsersController) Update() {
	fmt.Println("修改")
	c.TplName = "update.tpl"

}
func(c *UsersController) PostUpdate() {
	name := c.GetString("username")
	oldpassword := c.GetString("oldpassword")
	newpassword := c.GetString("newpassword")

	u := new(models.Userdata)
	u.Name = name 
	u.Password = oldpassword

	switch	flag := models.Update(u,newpassword); flag {
		case Success : c.Ctx.WriteString("修改成功")
		case Failed : c.Ctx.WriteString("修改密码失败")
		case Error : c.Ctx.WriteString("用户名或者旧密码错误")
	}

}

func (c *UsersController) Upload() {
	fmt.Println("上传")
	c.TplName = "upload.tpl"
}
func (c *UsersController) PostUpload() {
	fmt.Println("upload")
	f, h, err := c.GetFile("uploadname")
	defer f.Close()

	if err != nil {
		fmt.Println("getfile err", err)
	} else {
		c.SaveToFile("uploadname","static/upload/" + h.Filename)
	}

	c.Ctx.WriteString("上传成功")
}