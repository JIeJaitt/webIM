# 数据库增删改查一般套路
# 一、安装初始化
xorm.NewSession(driverName,dataSourceName)
# 二、定义实体
模型层model或者实体层entity
## 1、定义与表结构对应对象User
```go
type User struct {
    Id         int64     `xorm:"pk autoincr bigint(64)" form:"id" json:"id"`
    Mobile   string 		`xorm:"varchar(20)" form:"mobile" json:"mobile"`
    Passwd       string	`xorm:"varchar(40)" form:"passwd" json:"-"`   // 什么角色
    Avatar	   string 		`xorm:"varchar(150)" form:"avatar" json:"avatar"`
    Sex        string	`xorm:"varchar(2)" form:"sex" json:"sex"`   // 什么角色
    Nickname    string	`xorm:"varchar(20)" form:"nickname" json:"nickname"`   // 什么角色
    Salt       string	`xorm:"varchar(10)" form:"salt" json:"-"`   // 什么角色
    Online     int	`xorm:"int(10)" form:"online" json:"online"`   //是否在线
    Token      string	`xorm:"varchar(40)" form:"token" json:"token"`   // 什么角色
    Memo      string	`xorm:"varchar(140)" form:"memo" json:"memo"`   // 什么角色
    Createat   time.Time	`xorm:"datetime" form:"createat" json:"createat"`   // 什么角色
}
```
# 三、定义和业务相关的服务
服务层service,专门用来存放数据库业务服务的,如
注册、登录
## 2、查询单个用户Find,参数userId
```go
DbEngin.ID(userId).Get(&User)
```

## 3、查询满足某一类条件的Search
```go
result :=make([]User,0)
DbEngin.where("mobile=? ",moile).Find(&result)
DbEngin.where("mobile=? ",moile).Get(&User)
```

## 4、创建一条记录Create
```go
DBengin.InsertOne(&User)
```

## 5、修改某条记录Update
```go
DBengin.ID(userId).Update(&User)
// update ... where id = xx
DBengin.Where("a=? and b=?",a,b).Update(&User)
DBengin.Where("a=? and b=?",a,b).Cols("nick_name").Update(&User)
```

## 6、删除某条记录Delete
```go
DBengin.ID(userId).Delete(&User)
```

## 7、MD5加密函数
```go
import (
	"crypto/md5"
	"encoding/hex"
	"strings"
)

func Md5Encode(data string) string{
	h := md5.New()
	h.Write([]byte(data)) // 需要加密的字符串为 123456
	cipherStr := h.Sum(nil)

	return  hex.EncodeToString(cipherStr)

}
func MD5Encode(data string) string{
	return strings.ToUpper(Md5Encode(data))
}

func ValidatePasswd(plainpwd,salt,passwd string) bool{
	return Md5Encode(plainpwd+salt)==passwd
}
func MakePasswd(plainpwd,salt string) string{
	return Md5Encode(plainpwd+salt)
}
```     
# 四、控制器层调用
```go
var userServer server.UserServer
type UserCtrl struct{}

func (ctrl *UserCtrl)Register(w){
    user = userServer.Register(mobile,passwd)
}
```