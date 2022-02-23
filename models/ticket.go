package models

import (
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
)

type Ticket struct {
	Id   int64 `orm:"column(id)"`
	Stub rune  `orm:"column(stub)"`
}

// GetTicket 获取分号发布器id
// 在主键自增的情况下，分布式场景获取唯一id
func GetTicket() (id int64) {
	o := orm.NewOrm()
	var tik Ticket
	// replace into 刷新主键自增
	_, err := o.Raw("REPLACE INTO `t_ticket` (stub) VALUES (?);", 'a').Exec()
	if err != nil {
		logs.Error("replace into sql error")
		return 0
	}
	// 获取id
	err2 := o.QueryTable("t_ticket").Filter("stub", 'a').One(&tik)
	if err2 != nil {
		logs.Error("get ticket sql error")
		return 0
	}
	//logs.Info(tik)
	//logs.Info(tik.Id)
	return tik.Id
}
