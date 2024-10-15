package user

import (
	"blog-service/pkg/db"
	"blog-service/pkg/define"
	"blog-service/pkg/logger"
	"blog-service/pkg/models"

	"go.uber.org/zap"
)

func Register(u *models.User) (*define.UserInfo, error) {
	logger.Logger.Info("Register user", zap.Any("user", u))
	err := db.DB.Create(&u).Error
	if err != nil {
		logger.Logger.Error("Failed to create user", zap.Error(err))
		return nil, err
	}
	userInfo := define.UserInfo{
		ID:       u.ID,
		Username: u.Username,
		Email:    u.Email,
	}
	logger.Logger.Info("User created successfully", zap.Any("user_info", userInfo))
	return &userInfo, nil
}
