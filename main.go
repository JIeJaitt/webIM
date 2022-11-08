package main

import (
	"io"
	"net/http"
)

func main() {
	// 绑定请求和处理函数
	http.HandleFunc("/user/login",
		func(writer http.ResponseWriter,
			request *http.Request) {
			// 数据库操作
			// 逻辑处理

			// restapi json/xml返回

			// 1. 获取前端传递的参数
			// mobile,passwd
			// 解析参数
			// 如何获得参数

			// 解析参数 记得先 request.ParseForm()
			request.ParseForm()
			// 获取账号密码
			mobile := request.PostForm.Get("mobile")
			passwd := request.PostForm.Get("passwd")

			//
			loginok := false
			if mobile == "18600000000" && passwd == "123456" {
				loginok = true
			}

			str := `{"code":0,"data":{"id":1,"token":"test"}}`
			if !loginok {
				str = `{"code":-1,"msg":"密码不正确"}`
			}
			// 设置header 为 JSON 默认的text/html，所以特别指出返回的为application/json
			// 设置header
			writer.Header().Set("Content-Type", "application/json")
			// 设置200状态
			writer.WriteHeader(http.StatusOK)
			// 输出
			writer.Write([]byte(str))
			// 返回 json ok

			io.WriteString(writer, "hello,world!")
		})
	http.ListenAndServe(":8080", nil)
}
