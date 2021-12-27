package models

import (
	"time"
)

type RoomPeople struct {
	Pid        int       `json:"pid" xorm:"not null pk comment('人物ID') INT"`
	RoomId     int       `json:"room_id" xorm:"not null comment('房间ID') unique(room_role) INT"`
	RoleId     int       `json:"role_id" xorm:"not null comment('角色ID') unique(room_role) INT"`
	UserId     int       `json:"user_id" xorm:"not null comment('用户ID') INT"`
	CreateTime time.Time `json:"create_time" xorm:"not null comment('创建时间') DATETIME"`
	UpdateTime time.Time `json:"update_time" xorm:"not null comment('更新时间') DATETIME"`
}
