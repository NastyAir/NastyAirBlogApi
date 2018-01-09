package route

import (
	"net/http"
	"../service"
	"NastyAir/Blog/Utils"
)

func Login(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	//fmt.Println("method:", r.Method) //获取请求的方法
	var account, password []string
	if r.Method == "POST" {
		//请求的是登陆数据，那么执行登陆的逻辑判断
		account = r.Form["username"]
		password = r.Form["password"]
	}
	if len(account) < 1 || len(password) < 1 {
		Utils.CustomHttpCodeResponse("参数错误", http.StatusBadRequest, w)
	} else {
		var data = service.Login(account[0], password[0])
		Utils.JsonResponse(data, w)
	}
}
