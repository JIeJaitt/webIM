package model

import "time"

type Contact struct {
	Id       int64     `xorm:"pk autoincr bigint(20)" form:"id" json:"id"`
	Ownerid  int64     `xorm:"bigint(20)" form:"ownerid" json:"ownerid"`
	Dstobj   int64     `xorm:"bigint(20)" form:"dstobj" json:"dstobj"`
	Cate     int       `xorm:"int(11)" form:"cate" json:"cate"`
	Memo     string    `xorm:"varchar(120)" form:"memo" json:"memo"`
	Createat time.Time `xorm:"datetime" form:"createat" json:"createat"`
}

const (
	CONCAT_CATE_USER     = 0x01 //用户
	CONCAT_CATE_COMUNITY = 0x02 //群组
)
