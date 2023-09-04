package db

import (
	"dousheng/config"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os/user"
)

func Init() *gorm.DB {
	d := config.Db{
		User:     "root",
		Password: "zzh301029995663",
		DbName:   "dousheng",
		Ip:       "127.0.0.1",
		Port:     3306,
	}
	dsn := config.GetDsn(d)
	fmt.Println(dsn)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("数据库连接错误")
	}
	return db
}

// 创建
func Create() {
	db := Init()
	u := user.User{}
	db.AutoMigrate(&u)
}
