package main

import (
	_ "github.com/go-sql-driver/mysql"
	"html/template"
	"log"
	"net/http"
	"webIM/ctrl"
)

func RegisterView() {
	tpl, err := template.ParseGlob("view/**/*")
	if err != nil {
		// 模版渲染出错，直接打印并直接退出
		log.Fatal(err.Error())
	}
	for _, v := range tpl.Templates() {
		tplname := v.Name()
		http.HandleFunc(tplname, func(writer http.ResponseWriter, request *http.Request) {
			tpl.ExecuteTemplate(writer, tplname, nil)
		})
	}
}

func main() {
	// 绑定请求和处理函数
	http.HandleFunc("/user/login", ctrl.UserLogin)
	http.HandleFunc("/user/register", ctrl.UserRegister)
	http.HandleFunc("/contact/loadcommunity", ctrl.LoadCommunity)
	http.HandleFunc("/contact/loadfriend", ctrl.LoadFriend)
	http.HandleFunc("/contact/joincommunity", ctrl.JoinCommunity)
	http.HandleFunc("/contact/addfriend", ctrl.Addfriend)
	// 提供指定目录的静态文件支持
	http.Handle("/asset/", http.FileServer(http.Dir(".")))
	// 模版渲染
	RegisterView()
	// 启动Web服务器
	http.ListenAndServe(":8080", nil)
}
