/*
 * Copyright (c) 2017 - 2020, 奶爸<hi@nai.ba>
 * All rights reserved.
 */

package admin

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/naiba/nocd"
	"github.com/naiba/nocd/utils/mgin"
)

//User 用户管理
func User(us nocd.UserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		page := c.Query("page")
		var pageInt int64
		pageInt, _ = strconv.ParseInt(page, 10, 64)
		if pageInt < 0 {
			c.String(http.StatusForbidden, "GG")
			return
		}
		if pageInt == 0 {
			pageInt = 1
		}
		users, num := us.Users(pageInt-1, 20)
		c.HTML(http.StatusOK, "admin/user", mgin.CommonData(c, false, gin.H{
			"users":       users,
			"allPage":     num,
			"currentPage": pageInt,
		}))
	}
}

//UserToggle 用户状态管理
func UserToggle(us nocd.UserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		col := c.Param("col")
		act := c.Param("act")
		uid, err := strconv.ParseInt(id, 10, 64)
		if err != nil || uid < 1 {
			c.String(http.StatusForbidden, "ID有误，请重试")
			return
		}
		u, err := us.UserByGID(uid)
		if err != nil {
			c.String(http.StatusInternalServerError, "获取用户错误:"+err.Error())
			return
		}
		switch col {
		case "admin":
			u.IsAdmin = act == "on"
			err = us.Update(u)
			break
		case "block":
			u.IsBlocked = act == "on"
			err = us.Update(u)
			break
		}
		if err != nil {
			c.String(http.StatusInternalServerError, "更新用户错误:"+err.Error())
		}
	}
}
