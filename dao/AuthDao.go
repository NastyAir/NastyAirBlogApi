package dao

import (
	_ "github.com/go-sql-driver/mysql"
	"NastyAir/Blog/entity"
)

func UserFindByAccount(accountStr, pwd string) entity.UserCredentials {
	rows, err := Db.Query("SELECT account, password FROM user WHERE account=?", accountStr)
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
