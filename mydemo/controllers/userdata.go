package controllers

import (
	"github.com/astaxie/beego"
	_"github.com/go-sql-driver/mysql"
	"fmt"
	"sync"
	"time"

	"mygolang/mydemo/mydemo/models"


)

const (
	Success = iota
	Failed
	Error

)



var wg sync.WaitGroup
var lr LimitRate

type UsersController struct {
	beego.Controller
}

func init() {
	lr.SetRate(3)
	lr.lastAction = time.Now()
	
}

func (c *UsersController) Register() {
	wg.Add(1)
	go func() {
		if lr.Limit() {
			fmt.Println("注册")
			c.TplName = "register.tpl"
		}
		wg.Done()
	}()
	
	wg.Wait()
	
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
	wg.Add(1)
	go func() {
		if lr.Limit() {
			fmt.Println("登录")
			c.TplName = "login.tpl"
		}
		wg.Done()
	}()
	
	wg.Wait()


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
	//c.Ctx.WriteString("登陆成功")
	//c.TplName = "websocket.tpl"
	//c.TplName = "websocket.tpl"
	c.Redirect("/ws?uname="+name, 302)
}



func(c *UsersController) Update() {
	//fmt.Println("修改")
	//c.TplName = "update.tpl"

	wg.Add(1)
	go func() {
		if lr.Limit() {
			fmt.Println("修改")
			c.TplName = "update.tpl"
		}
		wg.Done()
	}()
	
	wg.Wait()

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
	//fmt.Println("上传")
	//c.TplName = "upload.tpl"

	wg.Add(1)
	go func() {
		if lr.Limit() {
			fmt.Println("上传")
			c.TplName = "upload.tpl"
		}
		wg.Done()
	}()
	
	wg.Wait()
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