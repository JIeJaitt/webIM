package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/user/login",
		func(writer http.ResponseWriter,
			request *http.Request) {
			request.ParseForm()

			mobile := request.PostForm.Get("mobile")
			passwd := request.PostForm.Get("passwd")

			loginok := false
			if mobile == "19118761673" && passwd == "123456" {
				loginok = true
			}

			str := `{"code":0,"data":{"id":1,"token":"test"}}`
			if !loginok {
				str = `{"code":-1,"msg":"密码不正确"}`
			}
			writer.Header().Set("Content-Type", "application/json")
			writer.WriteHeader(http.StatusOK)
			writer.Write([]byte(str))
		})
	http.ListenAndServe(":8080", nil)
}
