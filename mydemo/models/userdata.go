package models

import (
	"github.com/astaxie/beego/orm"
	"fmt"
)
//修改密码成功与否的枚举
const (
	Success = iota
	Failed
	Error

)
//数据表
type Userdata struct {
	Id int `orm:"auto"`
	Name string
	Password string
}

func init() {
	orm.RegisterModel(new(Userdata)) //注册数据表模型
	orm.RegisterDriver("mysql",orm.DRMySQL)//注册数据库驱动为mysql
	orm.RegisterDataBase("default","mysql","root:ilvylpo@/test?charset=utf8")
	orm.RegisterDataBase("test","mysql","root:ilvylpo@/test?charset=utf8")
}

//向数据库注册用户
func Register(u *Userdata) bool{
	orm.Debug = true
	b := true
	o := orm.NewOrm()
	o.Using("test")//使用的mysql数据库为test
	_, err := o.Insert(u)

	if err != nil {
		b = false
	}

	return b
}

func Login(u *Userdata)  bool{
	
	orm.Debug = true
	o := orm.NewOrm()
	o.Using("test")//使用的mysql数据库为test
	err := o.Read(u,"Name","Password")

	if err != nil {
		return false 
	}
	return  true
	
}


//修改密码
func Update(u *Userdata, password string) int{
	orm.Debug = true
	o := orm.NewOrm()
	o.Using("test")//使用的mysql数据库为test
	err := o.Read(u,"Name","Password")

	if err != nil {
		return Error
	}

	u.Password = password

	if num, err := o.Update(u); err != nil {
        fmt.Println(num)
        return Failed
    }
    return Success

}