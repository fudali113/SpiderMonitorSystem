package mysql

import (
	"fmt"
	"time"

	"github.com/astaxie/beego/orm"
)

type All struct {
	Id        int
	Pcid      string
	Ip        string
	Step      int
	Bid       string
	Sid       string
	All       string
	Exception string
	Time      time.Time `orm:"auto_now_add;type(datetime)"`
}

func (a *All) TableName() string {
	return "all_data"
}

type Finish struct {
	Id   int
	Pcid string
	Sid  string
	Bid  string
	Step int
}

type Exception struct {
	Id        int
	Pcid      string
	Ip        string
	Step      int
	Bid       string
	Exception string
	Data      string
	Time      time.Time `orm:"auto_now_add;type(datetime)"`
}

type HB struct {
	Id       int
	Pcid     string
	Ip       string
	Deadtime time.Time `orm:"auto_now_add;type(datetime)"`
}

func (a *HB) TableName() string {
	return "heartbeats"
}

func init() {
	fmt.Println("init db models")
	// 需要在init中注册定义的model
	orm.RegisterModel(new(All), new(Exception), new(HB), new(Finish))
}
