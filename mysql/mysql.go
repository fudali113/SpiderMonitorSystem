package mysql

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

var (
//db = orm.NewOrm().Using("default")
)

var (
	user   = beego.AppConfig.String("mysql.user")
	passwd = beego.AppConfig.String("mysql.passwd")
	host   = beego.AppConfig.String("mysql.host")
	port   = beego.AppConfig.String("mysql.port")
)

func init() {
	beego.Notice("init mysql conn")
	orm.RegisterDriver("mysql", orm.DRMySQL)
	mysqlConnStr := fmt.Sprintf("%s:%s@(%s:%s)/monitor?charset=utf8&loc=Local", user, passwd, host, port)
	orm.RegisterDataBase("default", "mysql", mysqlConnStr)
	//	orm.Debug = true
}

func InsertAll(all *All) int64 {
	db := orm.NewOrm()
	db.Using("default") // 默认使用 default，你可以指定为其他数据库
	r, e := db.Insert(all)
	beego.Notice("插入 all_data ")
	if e != nil {
		beego.Error(e)
		return 0
	}
	return r
}

func InsertExecption(all *Exception) int64 {
	db := orm.NewOrm()
	db.Using("default") // 默认使用 default，你可以指定为其他数据库
	r, e := db.Insert(all)
	beego.Notice("插入 exception")
	if e != nil {
		beego.Error(e)
		return 0
	}
	return r
}

func IOUFinish(all *Finish) int64 {
	db := orm.NewOrm()
	db.Using("default") // 默认使用 default，你可以指定为其他数据库
	r, e := db.InsertOrUpdate(all)
	beego.Notice("插入 finish")
	if e != nil {
		beego.Error(e)
		return 0
	}
	return r
}

func InsertHB(all *HB) int64 {
	db := orm.NewOrm()
	db.Using("default") // 默认使用 default，你可以指定为其他数据库
	r, e := db.Insert(all)
	beego.Notice("插入 heartbeat ")
	if e != nil {
		beego.Error(e)
		return 0
	}
	return r
}

func InsertTraffic(all *Traffic) int64 {
	db := orm.NewOrm()
	r, e := db.Insert(all)
	beego.Notice("插入 Traffic ")
	if e != nil {
		beego.Error(e)
		return 0
	}
	return r
}

func InsertCS(all *CompStatus) int64 {
	db := orm.NewOrm()
	r, e := db.Insert(all)
	beego.Notice("插入 computer status ")
	if e != nil {
		beego.Error(e)
		return 0
	}
	return r
}
