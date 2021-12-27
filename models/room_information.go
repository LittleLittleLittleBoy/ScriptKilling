package models

import (
	"time"
)

type RoomInformation struct {
	Iid        int       `json:"iid" xorm:"not null pk autoincr comment('房间信息ID') INT"`
	InfoId     int       `json:"info_id" xorm:"not null comment('信息ID') unique(room_infor) INT"`
	RoomId     int       `json:"room_id" xorm:"not null comment('房间ID') unique(room_infor) INT"`
	UserId     int       `json:"user_id" xorm:"not null comment('角色ID') INT"`
	CreateTime time.Time `json:"create_time" xorm:"not null comment('创建时间') DATETIME"`
	UpdateTime time.Time `json:"update_time" xorm:"not null comment('更新时间') DATETIME"`
}
