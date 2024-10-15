package user

import (
	"blog-service/pkg/db"
	"blog-service/pkg/logger"
	"blog-service/pkg/models"

	"go.uber.org/zap"
)

func Register(u *models.User) error {
	logger.Logger.Info("Register user", zap.Any("user", u))
	err := db.DB.Create(&u).Error
	if err != nil {
		logger.Logger.Error("Failed to create user", zap.Error(err))
		return err
	}
	logger.Logger.Info("User created successfully")
	return nil
}
