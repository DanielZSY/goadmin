package model

import (
	"time"
)

type BUser struct {
	Id         int64         `json:"uid" form:"uid" gorm:"primary_key"`
	Username   string        `json:"username" form:"username"`
	Password   string        `json:"password" form:"password"`
	Gender     int           `json:"gender" form:"gender"` // 1=男 2=女
	Enable     int           `json:"enable" form:"enable"` // 1 启用 0 禁用
	Name       string        `json:"name" form:"name"`
	Age        int           `json:"age" form:"age"`
	Phone      string        `json:"phone" form:"phone"`
	Email      string        `json:"email" form:"email"`
	Avatar     string        `json:"avatar" form:"avatar"`
	CreateTime time.Time     `json:"createTime" form:"createTime"`
	Online     bool          `json:"online" gorm:"-"` // 是否在线
}
