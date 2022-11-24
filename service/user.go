package service

import (
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"math/rand"
	"time"
	"webIM/model"
	"webIM/util"
)

type UserService struct {
}

// 注册函数
func (s *UserService) Register(mobile, plainpwd, nickname, avatar, sex string) (user model.User, err error) {
	// 检测手机号是否存在
	tmp := model.User{}
	_, err = DbEngin.Where("mobile=? ", mobile).Get(&tmp)
	if err != nil {
		return tmp, err
	}
	// 如果存在则返回提示已经注册
	if tmp.Id > 0 {
		return tmp, errors.New("该手机号已经注册")
	}
	// 否则拼接插入数据
	tmp.Mobile = mobile
	tmp.Avatar = avatar
	tmp.Nickname = nickname
	tmp.Sex = sex

	tmp.Salt = fmt.Sprintf("%06d", rand.Int31n(10000))
	tmp.Passwd = util.MakePasswd(plainpwd, tmp.Salt)

	tmp.Createat = time.Now()
	// 返回新用户的信息
	_, err = DbEngin.InsertOne(&tmp)
	return tmp, nil
}

// 登陆函数
func (s *UserService) Login(mobile, plainpwd string) (user model.User, err error) {
	// 通过手机号查询用户
	tmp := model.User{}
	DbEngin.Where("mobile=?", mobile).Get(&tmp)
	// 如果没有找到
	if tmp.Id == 0 {
		return tmp, errors.New("该用户不存在")
	}
	// 查询到手机号之后比对密码
	if util.ValidatePasswd(plainpwd, tmp.Salt, tmp.Passwd) {
		return tmp, errors.New("密码不正确")
	}
	// 为了安全起见刷新 token
	str := fmt.Sprintf("%d", time.Now().Unix())
	token := util.MD5Encode(str)
	tmp.Token = token
	DbEngin.ID(tmp.Id).Cols("token").Update(&tmp)

	return tmp, nil
}
