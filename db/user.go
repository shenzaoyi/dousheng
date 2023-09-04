package db

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name             string
	Follow_count     int
	Follower_count   int
	Is_follow        bool
	Avatar           string
	Background_image string
	Signature        string
	Total_favorited  int
	Work_count       int
	Favorite_count   int
	//ID               int
}

func AddUser(u *User) string {
	db := Init()
	result := db.Create(u)
	if result.Error != nil {
		return "添加用户信息错误"
	}
	return ""
}
