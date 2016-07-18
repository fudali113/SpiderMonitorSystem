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
	fmt.Println("init mysql conn")
	orm.RegisterDriver("mysql", orm.DRMySQL)
	mysqlConnStr := fmt.Sprintf("%s:%s@(%s:%s)/monitor?charset=utf8", user, passwd, host, port)
	fmt.Println(mysqlConnStr)
	orm.RegisterDataBase("default", "mysql", mysqlConnStr)
	orm.Debug = true
	fmt.Println("init mysql conn end")
}

func InsertAll(all *All) int64 {
	db := orm.NewOrm()
	db.Using("default") // 默认使用 default，你可以指定为其他数据库
	r, e := db.Insert(all)
	if e != nil {

		fmt.Println(e)
		return 0
	}
	fmt.Println(r)
	return 1
}

func InsertExecption(all *Exception) int64 {
	db := orm.NewOrm()
	db.Using("default") // 默认使用 default，你可以指定为其他数据库
	_, e := db.Insert(all)
	if e != nil {
		fmt.Println(e)
		return 0
	}
	return 1
}

func IOUFinish(all *Finish) int64 {
	db := orm.NewOrm()
	db.Using("default") // 默认使用 default，你可以指定为其他数据库
	r, e := db.InsertOrUpdate(all)
	if e != nil {
		fmt.Println(e)
		return 0
	}
	fmt.Println(r)
	return r
}

func InsertHB(all *HB) int64 {
	db := orm.NewOrm()
	db.Using("default") // 默认使用 default，你可以指定为其他数据库
	_, e := db.Insert(all)
	if e != nil {
		fmt.Println(e)
		return 0
	}
	return 1
}
