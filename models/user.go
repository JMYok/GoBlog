package models

import "time"

type User struct {
	Uid      int       `json:"uid"`
	UserName string    `json:"user_name"`
	PassWd   string    `json:"pass_wd"`
	Avatar   string    `json:"avatar"`
	CreateAt time.Time `json:"create_at"`
	UpdateAt time.Time `json:"update_at"`
}

type UserInfo struct {
	Uid      int    `json:"uid"`
	UserName string `json:"user_name"`
	Avatar   string `json:"avatar"`
}
