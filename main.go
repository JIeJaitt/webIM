package main

import (
	"encoding/json"
	"html/template"
	"io"
	"log"
	"net/http"
)

func userLogin(writer http.ResponseWriter, request *http.Request) {
	// 解析参数
	request.ParseForm()
	// 获取账号密码
	mobile := request.PostForm.Get("mobile")
	passwd := request.PostForm.Get("passwd")

	loginok := false
	if mobile == "18600000000" && passwd == "123456" {
		loginok = true
	}
	if loginok {
		data := make(map[string]interface{})
		data["id"] = 1
		data["token"] = "test"
		Resp(writer, 0, data, "")
	} else {
		Resp(writer, -1, nil, "密码不正确")
	}
	io.WriteString(writer, "hello,world!")
}

type H struct {
	Code int         `json:"code""`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

func Resp(w http.ResponseWriter, code int, data interface{}, msg string) {
	w.Header().Set("Content-Type", "application/json")
	// 设置200状态
	w.WriteHeader(http.StatusOK)
	// 定义一个结构体
	h := H{
		Code: code,
		Msg:  msg,
		Data: data,
	}
	// 将结构体转化为字符串
	ret, err := json.Marshal(h)
	if err != nil {
		log.Println(err.Error())
	}
	// 输出
	w.Write(ret)
}

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
	http.HandleFunc("/user/login", userLogin)
	// 提供指定目录的静态文件支持
	http.Handle("/asset/", http.FileServer(http.Dir(".")))
	// 模版渲染
	RegisterView()
	// 启动Web服务器
	http.ListenAndServe(":8080", nil)
}
