package models

import (
	"time"
)

type ScriptInfo struct {
	Sid        int       `json:"sid" xorm:"not null pk autoincr comment('剧本 ID') INT"`
	Title      string    `json:"title" xorm:"not null comment('剧本名称') unique VARCHAR(255)"`
	RoleNum    int       `json:"role_num" xorm:"not null comment('人数') INT(10)"`
	Cover      string    `json:"cover" xorm:"not null comment('封面') VARCHAR(255)"`
	CreateTime time.Time `json:"create_time" xorm:"not null comment('创建时间') DATETIME"`
	UpdateTime time.Time `json:"update_time" xorm:"not null comment('更新时间') DATETIME"`
}
