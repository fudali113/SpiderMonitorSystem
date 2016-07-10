package mysql

import (
	"fmt"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

var (
//db = orm.NewOrm().Using("default")
)

func init() {
	fmt.Println("init mysql conn")
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:Tt7896357@(120.26.235.113:6666)/monitor?charset=utf8")
	orm.Debug = true
}

func InsertAll(all *All) int64 {
	db := orm.NewOrm()
	db.Using("default") // 默认使用 default，你可以指定为其他数据库
	_, e := db.Insert(all)
	if e != nil {
		fmt.Println(e)
		return 0
	}
	return 1
}

func InsertExecption(all *Execption) int64 {
	db := orm.NewOrm()
	db.Using("default") // 默认使用 default，你可以指定为其他数据库
	_, e := db.Insert(all)
	if e != nil {
		fmt.Println(e)
		return 0
	}
	return 1
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
