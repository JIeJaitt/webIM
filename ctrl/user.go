package ctrl

import (
	"fmt"
	"math/rand"
	"net/http"
	"webIM/model"
	"webIM/service"
	"webIM/util"
)

func UserLogin(writer http.ResponseWriter, request *http.Request) {
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
		util.RespOk(writer, 0, "")
	} else {
		util.RespFail(writer, "密码不正确")
	}
}

var userService service.UserService

func UserRegister(writer http.ResponseWriter, request *http.Request) {
	request.ParseForm()
	mobile := request.PostForm.Get("mobile")
	plainpwd := request.PostForm.Get("passwd")
	nickname := fmt.Sprintf("user%06d", rand.Int31())
	avatar := ""
	sex := model.SEX_UNKNOW

	user, err := userService.Register(mobile, plainpwd, nickname, avatar, sex)
	if err != nil {
		util.RespFail(writer, err.Error())
	} else {
		util.RespOk(writer, user, "")
	}

}
