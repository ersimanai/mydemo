package controllers

import (
	"github.com/astaxie/beego"
	"sync"
	"time"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["register"] = "register"
	c.TplName = "index.tpl"
}



type LimitRate struct {
	rate int 
	interval time.Duration
	lastAction time.Time 
	lock sync.Mutex
}

//Limit 限速
func (l LimitRate) Limit() bool {
	result := false

	for {
		l.lock.Lock()

		if  time.Now().Sub(l.lastAction) > l.interval {
			l.lastAction = time.Now()
			result = true
		}
		l.lock.Unlock()

		if result {
			return result
		}

		time.Sleep(l.interval)
	}

	return result
}
//设置rate
func (l *LimitRate) SetRate(r  int) {
	l.rate = r 
	l.interval = time.Microsecond * time.Duration(1000*1000/l.rate)
}
//获取rate
func (l *LimitRate) GetRate() int {
	return l.rate
}