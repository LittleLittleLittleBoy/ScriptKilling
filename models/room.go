package models

import (
	"time"
)

type Room struct {
	Rid        int       `json:"rid" xorm:"not null pk autoincr comment('房间ID') INT"`
	Sid        int       `json:"sid" xorm:"not null comment('剧本ID') INT"`
	State      int       `json:"state" xorm:"not null comment('房间状态') INT"`
	CreateTime time.Time `json:"create_time" xorm:"not null comment('创建时间') DATETIME"`
	UpdateTime time.Time `json:"update_time" xorm:"not null comment('更新时间') DATETIME"`
}
