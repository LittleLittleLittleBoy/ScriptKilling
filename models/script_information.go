package models

import (
	"time"
)

type ScriptInformation struct {
	Iid        int       `json:"iid" xorm:"not null pk autoincr comment('信息ID') INT"`
	Type       int       `json:"type" xorm:"not null comment('信息类型') INT(255)"`
	Title      string    `json:"title" xorm:"not null default '' comment('信息名称') VARCHAR(255)"`
	Stage      int       `json:"stage" xorm:"not null comment('所属的阶段') INT(255)"`
	Url        string    `json:"url" xorm:"not null comment('信息url') VARCHAR(255)"`
	Uid        int       `json:"uid" xorm:"not null comment('信息归属的角色ID') INT"`
	Sid        int       `json:"sid" xorm:"not null comment('信息归属的剧本ID') INT"`
	CreateTime time.Time `json:"create_time" xorm:"not null comment('创建时间') DATETIME"`
	UpdateTime time.Time `json:"update_time" xorm:"not null comment('修改时间') DATETIME"`
}
