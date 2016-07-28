package mysql

import (
	"fmt"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type TimeFilter struct {
	Start string
	End   string
}

func (t *TimeFilter) check() (res bool) {
	if t.Start != "" && t.End != "" {
		res = true
	} else {
		res = false
	}
	return
}

type StepExecAllRatio struct {
	Step   int `json:"step"`
	Normal int `json:"normal"`
	Exec   int `json:"exec"`
}

func GetExecAllRatio(t *TimeFilter) []StepExecAllRatio {
	db := orm.NewOrm()
	var sears []StepExecAllRatio
	sqlStr := ""
	fmt.Println(t.check())
	if !t.check() {
		sqlStr = "select step, count(step) normal , count(case when exception != '' then step end) exec  FROM all_data GROUP BY step"
	} else {
		sqlStr = fmt.Sprintf("select step, count(step) normal , count(case when exception != '' then step end) exec  FROM all_data Where time >= '%s' AND time <= '%s' GROUP BY step", t.Start, t.End)
	}
	_, err := db.Raw(sqlStr).QueryRows(&sears)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	fmt.Println(sears)
	return sears
}

func GetStepFinish(t *TimeFilter) []StepExecAllRatio {
	db := orm.NewOrm()
	var sears []StepExecAllRatio
	sqlStr := ""
	if !t.check() {
		sqlStr = "SELECT step , count(step) normal from finish GROUP BY step"
	} else {
		sqlStr = fmt.Sprintf("SELECT step , count(step) normal from finish  Where time >= '%s' AND time <= '%s' GROUP BY step", t.Start, t.End)
	}
	_, err := db.Raw(sqlStr).QueryRows(&sears)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	fmt.Println(sears)
	return sears
}

type PcDownRatio struct {
	Pcid  string `json:"pcid"`
	Count int    `json:"count"`
}

func GetPcDownRatio(t *TimeFilter) []PcDownRatio {
	db := orm.NewOrm()
	var sears []PcDownRatio
	sqlStr := ""
	if !t.check() {
		sqlStr = "SELECT pcid , count(pcid) count from heartbeat GROUP BY pcid"
	} else {
		sqlStr = fmt.Sprintf("SELECT pcid , count(pcid) count from heartbeat Where deadtime >= '%s' AND deadtime <= '%s'  GROUP BY pcid ", t.Start, t.End)
	}
	_, err := db.Raw(sqlStr).QueryRows(&sears)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	fmt.Println(sears)
	return sears
}
