package models

import (
	"time"
)

type ScriptRole struct {
	Rid        int       `json:"rid" xorm:"not null pk autoincr comment('角色ID') INT"`
	Name       string    `json:"name" xorm:"not null comment('姓名') VARCHAR(255)"`
	Sid        int       `json:"sid" xorm:"not null comment('剧本ID') INT"`
	Cover      string    `json:"cover" xorm:"not null comment('角色封面') VARCHAR(255)"`
	CreateTime time.Time `json:"create_time" xorm:"not null comment('创建时间') DATETIME"`
	UpdateTime time.Time `json:"update_time" xorm:"not null comment('更新时间') DATETIME"`
}
