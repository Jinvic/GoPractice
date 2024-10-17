package user

import (
	"blog-service/pkg/db"
	"blog-service/pkg/define"
	"blog-service/pkg/logger"
	"blog-service/pkg/models"
	"errors"

	"go.uber.org/zap"
)

func Register(u *models.User) (*define.UserInfo, error) {
	logger.Logger.Info("Register user", zap.Any("user", u))
	err := db.DB.Create(&u).Error
	if err != nil {
		return nil, err
	}
	userInfo := define.UserInfo{
		ID:       u.ID,
		Username: u.Username,
		Email:    u.Email,
	}
	return &userInfo, nil
}

func Login(username string, password string) (*define.UserInfo, error) {
	logger.Logger.Info("Login user", zap.Any("user", username))
	user := models.User{}
	err := db.DB.Where("username = ?", username).First(&user).Error
	if err != nil {
		return nil, err
	}
	if user.Password != password {
		return nil, errors.New("password not match")
	}
	userInfo := define.UserInfo{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
	}
	return &userInfo, nil
}

func Delete(userInfo *define.UserInfo) error {
	logger.Logger.Info("Delete user", zap.Any("user", userInfo))
	err := db.DB.Where("id = ?", userInfo.ID).Unscoped().Delete(&models.User{}).Error
	if err != nil {
		return err
	}
	return nil
}

func List() (*[]define.UserInfo, error) {
	logger.Logger.Info("List users")
	users := []models.User{}
	err := db.DB.Find(&users).Error
	if err != nil {
		return nil, err
	}
	userInfos := []define.UserInfo{}
	for _, user := range users {
		userInfos = append(userInfos, define.UserInfo{
			ID:       user.ID,
			Username: user.Username,
			Email:    user.Email,
		})
	}
	return &userInfos, nil
}
