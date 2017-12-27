package route

import (
	"net/http"
	"fmt"
	"../dao"
	"../common"
	"encoding/json"
)

func Login(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	//fmt.Println("method:", r.Method) //获取请求的方法
	var account, password []string
	if r.Method == "POST" {
		//请求的是登陆数据，那么执行登陆的逻辑判断
		account = r.Form["username"]
		password = r.Form["password"]
		//fmt.Println("username:", account)
		//fmt.Println("password:", password)
	}
	var isAuth bool
	isAuth = dao.UserFindByAccount(account[0], password[0])
	msg := new(common.RestMsg)
	var data []byte
	if isAuth {
		msg.Code = common.SUCCESS
		data, _ = json.Marshal(msg)
	} else {
		msg.Code = common.FAIL
		data, _ = json.Marshal(msg)
	}
	fmt.Fprintf(w, string(data))
}
