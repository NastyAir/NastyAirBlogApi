package route

import (
	"net/http"
	"fmt"
	"../service"
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
	var msg = service.Login(account[0], password[0])
	var data []byte
	data, _ = json.Marshal(msg)
	fmt.Fprintf(w, string(data))
}
