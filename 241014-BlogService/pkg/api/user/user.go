package user

import (
	"blog-service/pkg/db"
	"blog-service/pkg/models"
	"time"

	"github.com/spf13/viper"
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