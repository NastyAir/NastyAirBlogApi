package dao

import "database/sql"

var Db *sql.DB

func init() {
	Db, _ = sql.Open("mysql", "root:@/nablog?charset=utf8")
	Db.SetMaxOpenConns(2000)
	Db.SetMaxIdleConns(1000)
	Db.Ping()
}
