package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
	Email    string `gorm:"default:no email"`
}

var db *gorm.DB

func init() {
	dsn := "root:root@tcp(127.0.0.1:3306)/userregisterlogin?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&User{})
}

func login(username string, password string) (user *User, errmsg string) {

	var cnt int64
	db.Model(&User{}).Where("username = ?", username).Count(&cnt)
	if cnt == 0 {
		errmsg = "用户不存在"
		return
	}

	err := db.Where("username = ?", username).First(&user).Error
	if err != nil {
		errmsg = err.Error()
	} else {
		errmsg = ""
	}

	if password != user.Password {
		errmsg = "用户名或密码错误"
	}

	return user, errmsg
}

func register(user *User) (errmsg string) {

	var cnt int64
	db.Where("username = ?", user.Username).Count(&cnt)
	if cnt != 0 {
		errmsg = "用户已存在"
		return
	}

	err := db.Create(&user).Error

	if err != nil {
		errmsg = err.Error()
	} else {
		errmsg = ""
	}
	return errmsg
}

func getUser(userID uint) (user *User) {
	db.First(&user, userID)
	return
}
