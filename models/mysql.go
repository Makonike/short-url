package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	logs.Info("mysql init...")
	user := beego.AppConfig.String("mysqluser")
	pass := beego.AppConfig.String("mysqlpass")
	url := beego.AppConfig.String("mysqlurl")
	db := beego.AppConfig.String("mysqldb")
	conn := user + ":" + pass + "@tcp(" + url + ")/" + db + "?charset=utf8"
	err := orm.RegisterDataBase("default", "mysql", conn, 30)
	if err != nil {
		logs.Error("mysql init error")
		return
	}
	orm.RegisterModelWithPrefix("t_", new(Ticket))
	orm.RegisterModelWithPrefix("t_", new(Short))
	orm.Debug = true
	err2 := orm.RunSyncdb("default", false, true)
	if err2 != nil {
		logs.Error("db run sync error")
		return
	}
	logs.Info("mysql init finish...")
}
