package user

import (
	"blog-service/pkg/db"
	"blog-service/pkg/logger"
	"blog-service/pkg/models"
	"blog-service/pkg/services/auth"
	"time"

	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func checkDuplicateUsername(username string) (bool, error) {
	var count int64
	err := db.DB.Model(&models.User{}).Where("username = ?", username).Count(&count).Error
	return count > 0, err
}

func checkDuplicateEmail(email string) (bool, error) {
	var count int64
	err := db.DB.Model(&models.User{}).Where("email = ?", email).Count(&count).Error
	return count > 0, err
}

func getExpiredAt() time.Time {
	expiredHour := viper.GetInt("token.expired.hour")
	expiredMinute := viper.GetInt("token.expired.minute")
	expiredSecond := viper.GetInt("token.expired.second")
	return time.Now().Add(time.Hour*time.Duration(expiredHour) +
		time.Minute*time.Duration(expiredMinute) +
		time.Second*time.Duration(expiredSecond))
}


func banOldToken(uid uint) error {
	// check if token already exists
	if hasToken, err := auth.HasToken(uid); err != nil {
		logger.Logger.Error("Failed to check if token exists", zap.Any("position", "api"), zap.Error(err))
		return err
	} else if hasToken {
		// ban old token
		oldToken, err := auth.GetToken(uid)
		if err != nil {
			logger.Logger.Error("Failed to get old token", zap.Any("position", "api"), zap.Error(err))
			return err
		}

		claims, err := auth.ParseToken(oldToken)
		if err != nil {
			logger.Logger.Error("Failed to parse old token", zap.Any("position", "api"), zap.Error(err))
			return err
		}

		err = auth.BanToken(oldToken, claims.RegisteredClaims.ExpiresAt.Time)
		if err != nil {
			logger.Logger.Error("Failed to ban old token", zap.Any("position", "api"), zap.Error(err))
			return err
		}
	}

	return nil
}