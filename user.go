/*
 * Copyright (c) 2017 - 2020, 奶爸<hi@nai.ba>
 * All rights reserved.
 */

package nocd

import "github.com/jinzhu/gorm"

//User 用户
type User struct {
	gorm.Model
	//用户GitHubID
	GID        uint `gorm:"unique_index"`
	GName      string
	GLogin     string
	Pubkey     string
	PrivateKey string
	//Server酱推送Key
	Sckey        string       `form:"sckey" binding:"alphanum,min=20"`
	PushSuccess  bool         `form:"push_success" binding:"exists"`
	Servers      []Server     `form:"-"`
	Repositories []Repository `form:"-"`
	Pipelines    []Pipeline   `form:"-"`
	IsBlocked    bool
	IsAdmin      bool
	// 用户Token
	Token string
}

//UserService 用户服务
type UserService interface {
	UserByGID(gid int64) (*User, error)
	Create(u *User) error
	Update(u *User) error
	Verify(uid, token string) (*User, error)
	Users(page, size int64) ([]*User, int64)
}
