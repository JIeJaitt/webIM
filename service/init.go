package service

import (
	"errors"
	"fmt"
	"github.com/go-xorm/xorm"
	"log"
	"webIM/model"
)

var DbEngin *xorm.Engine

func init() {
	drivename := "mysql"
	DsName := "root:12345678@(127.0.0.1:3306)/chat?charset=utf8"
	err := errors.New("")
	DbEngin, err = xorm.NewEngine(drivename, DsName)
	if err != nil && "" != err.Error() {
		log.Fatal(err.Error())
	}
	// 是否显示SQL
	DbEngin.ShowSQL(true)
	// 数据库最大打开的连接数
	DbEngin.SetMaxOpenConns(2)
	// 自动创建表结构
	DbEngin.Sync2(new(model.User),
		new(model.Contact),
		new(model.Community))
	// DbEngin = dbengin
	fmt.Println("init data base ok")
}
