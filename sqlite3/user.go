/*
 * Copyright (c) 2017 - 2020, 奶爸<hi@nai.ba>
 * All rights reserved.
 */

package sqlite3

import (
	"github.com/jinzhu/gorm"
	"github.com/naiba/nocd"
)

//UserService 用户服务
type UserService struct {
	DB *gorm.DB
}

//UserByGID 根据GitHubID获取用户
func (us *UserService) UserByGID(gid int64) (*nocd.User, error) {
	var u nocd.User
	return &u, us.DB.Where("g_id = ?", gid).First(&u).Error
}

//UserByGName 根据 GitHubName 获取用户
func (us *UserService) UserByGName(name string) (*nocd.User, error) {
	var u nocd.User
	return &u, us.DB.Where("g_name = ?", name).First(&u).Error
}

//Users 获取所有用户
func (us *UserService) Users(page, size int64) ([]*nocd.User, int64) {
	var ul []*nocd.User
	var num int64
	us.DB.Offset(page * size).Limit(size).Order("updated_at DESC").Find(&ul)
	for _, u := range ul {
		us.DB.Model(&u).Select("id").Related(&u.Servers)
		us.DB.Model(&u).Select("id").Related(&u.Repositories)
		us.DB.Model(&u).Select("id").Related(&u.Pipelines)
	}
	us.DB.Model(&nocd.User{}).Count(&num)
	if num%size == 0 {
		num = num / size
	} else {
		num = num/size + 1
	}
	return ul, num
}

//Create 创建用户
func (us *UserService) Create(u *nocd.User) error {
	return us.DB.Create(u).Error
}

//Update 更新用户
func (us *UserService) Update(u *nocd.User) error {
	return us.DB.Save(u).Error
}

//Verify 校验用户
func (us *UserService) Verify(uid, token string) (*nocd.User, error) {
	var u nocd.User
	return &u, us.DB.Where("id = ? AND token = ?", uid, token).First(&u).Error
}

// Transfer 账户迁移
func (us *UserService) Transfer(origin *nocd.User, dist *nocd.User) error {
	tx := us.DB.Begin()
	tx.Table("servers").Where("user_id = ?", origin.ID).Updates(map[string]interface{}{
		"user_id": dist.ID,
	})
	tx.Table("repositories").Where("user_id = ?", origin.ID).Updates(map[string]interface{}{
		"user_id": dist.ID,
	})
	tx.Table("pipelines").Where("user_id = ?", origin.ID).Updates(map[string]interface{}{
		"user_id": dist.ID,
	})
	return tx.Commit().Error
}
