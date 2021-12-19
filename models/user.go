package models

import (
	"time"
)

type User struct {
	Uid        int       `json:"uid" xorm:"not null pk autoincr comment('Primary Key') INT"`
	Username   string    `json:"username" xorm:"not null comment('User Name') unique VARCHAR(255)"`
	Password   string    `json:"password" xorm:"not null comment('Password') VARCHAR(255)"`
	CreateTime time.Time `json:"create_time" xorm:"comment('Create Time') DATETIME"`
	UpdateTime time.Time `json:"update_time" xorm:"comment('Update Time') DATETIME"`
}
