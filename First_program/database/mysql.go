package database

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
	logging "github.com/op/go-logging"
)

var SQLdb *sql.DB
var logs = logging.MustGetLogger("example")

func Opensql() {
	var err error
	SQLdb, err = sql.Open("mysql", "root:zhouqifan@tcp(127.0.0.1:3306)/articles?charset=utf8")
	if err != nil {
		log.Fatal(err.Error())
	}
	err = SQLdb.Ping()
	if err != nil {
		log.Fatal(err.Error())
	}
	logs.Debugf("数据库连接成功")
}
