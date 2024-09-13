package main

import (
	"errors"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `form:"username"`
	Password string `form:"password"`
}

var db *gorm.DB

func init() {
	//Data Source Name 数据源名称
	dsn := "root:root@tcp(127.0.0.1:3306)/db240909?charset=utf8mb4&parseTime=True&loc=Local&&timeout=10s"
	//链接MySQL数据库
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database: " + err.Error())
	}
	//自动迁移，根据模型生成或更新数据表
	db.AutoMigrate(&User{})
}

func regist(username, password string) (uint, error) {
	var user User
	result := db.Where("username = ?", username).First(&user)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		user.Username = username
		user.Password = password
		db.Create(&user)
		db.Where("username = ?", username).Find(&user)
		return user.ID, nil
	} else if user.Username == username {
		return 0, errors.New("用户名已被注册")
	} else {
		return 0, result.Error
	}
}

func login(username, password string) (uint, error) {
	var user User
	result := db.Where("username = ?", username).First(&user)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return 0, errors.New("用户名不存在")
	}
	if user.Password != password {
		return 0, errors.New("用户名和密码不匹配")
	}
	return user.ID, nil
}

func userInfo(userid uint) (string, string, error) {
	var user User
	result := db.Find(&user, userid)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return "", "", errors.New("用户不存在")
	}
	return user.Username, user.Password, nil
}
