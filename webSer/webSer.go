package webSer

import "net/http"
import (
	"../route"
	"log"
)

func Run() {
	//getConfig()
	registerRoute()
	setListenAndServe(":9090")
}

func setListenAndServe(addr string) {
	err := http.ListenAndServe(addr, nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func registerRoute() {
	//设置静态资源访问
	http.Handle("/dist/", http.StripPrefix("/dist/", http.FileServer(http.Dir("./dist"))))
	http.HandleFunc("/", route.IndexPage)
	http.HandleFunc("/login", route.Login) //设置访问的路由
}
