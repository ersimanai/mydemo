package models

import (
	"github.com/astaxie/beego/orm"
	_"fmt"
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
func Register(u *Userdata) {
	orm.Debug = true
	o := orm.NewOrm()
	o.Using("test")//使用的mysql数据库为test
	o.Insert(u)
}
//修改密码
func Update(u *Userdata) {
	orm.Debug = true
	o := orm.NewOrm()
	o.Using("test")//使用的mysql数据库为test
	o.Update(u)
}