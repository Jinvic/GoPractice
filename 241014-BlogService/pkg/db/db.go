package db

import (
	"blog-service/pkg/logger"
	"blog-service/pkg/models"
	"fmt"

	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		viper.GetString("database.user"),
		viper.GetString("database.password"),
		viper.GetString("database.host"),
		viper.GetInt("database.port"),
		viper.GetString("database.database"),
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Logger.Error("Error connecting to database", zap.Error(err))
		panic(err)
	}
	logger.Logger.Info("Successfully connected to database")

	err = db.AutoMigrate(&models.User{}, &models.Article{})
	if err != nil {
		logger.Logger.Error("Error migrating database", zap.Error(err))
		panic(err)
	}
	logger.Logger.Info("Successfully migrated database")

	DB = db
}
