package mysql

import (
	"fmt"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type StepExecAllRatio struct {
	Step   int `json:"step"`
	Normal int `json:"normal"`
	Exec   int `json:"exec"`
}

func GetExecAllRatio() []StepExecAllRatio {
	db := orm.NewOrm()
	var sears []StepExecAllRatio
	_, err := db.Raw("select step, count(step) normal , count(case when execption != '' then step end) exec  FROM all_data GROUP BY step").QueryRows(&sears)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	fmt.Println(sears)
	return sears
}

func GetStepFinish() []StepExecAllRatio {
	db := orm.NewOrm()
	var sears []StepExecAllRatio
	_, err := db.Raw("SELECT step , count(step) normal from finish GROUP BY step").QueryRows(&sears)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	fmt.Println(sears)
	return sears
}
