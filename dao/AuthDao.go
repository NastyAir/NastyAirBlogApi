package dao

import (
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"NastyAir/Blog/entity"
)

func UserFindByAccount(accountStr, pwd string) entity.UserCredentials {
	db, err := sql.Open("mysql", "root:@/nablog?charset=utf8")
	defer db.Close()
	checkErr(err)
	rows, err := db.Query("SELECT account, password FROM user WHERE account=?", accountStr)
	checkErr(err)
	//var userid, account, name, password, nickname, tel, email, avatar, place, creator, updater, description string
	var user entity.UserCredentials
	for rows.Next() {
		err = rows.Scan(&user.Username, &user.Password)
	}
	defer rows.Close()
	checkErr(err)
	return user
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
