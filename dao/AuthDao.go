package dao

import (
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	//"fmt"
)

func UserFindByAccount(accountStr, pwd string) bool {
	db, err := sql.Open("mysql", "root:@/nablog?charset=utf8")
	defer db.Close()
	checkErr(err)
	rows, err := db.Query("SELECT account, password FROM user WHERE account=?", accountStr)
	checkErr(err)
	//var id, status, gender int
	//var createtime, updatetime, birthday string
	//var userid, account, name, password, nickname, tel, email, avatar, place, creator, updater, description string
	var account, password string
	for rows.Next() {
		err = rows.Scan(&account, &password)
	}
	defer rows.Close()
	checkErr(err)
	//fmt.Println(account)
	//fmt.Println(password)
	isAuth := false
	isAuth = pwd == password
	return isAuth
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
