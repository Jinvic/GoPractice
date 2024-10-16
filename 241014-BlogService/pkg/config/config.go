package config

import (
	"blog-service/pkg/logger"

	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func InitConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath("./pkg/config")
	err := viper.ReadInConfig()
	if err != nil {
		logger.Logger.Error("Error reading config file", zap.Any("position", "init"), zap.Error(err))
		panic(err)
	}
	logger.Logger.Info("Successfully read config file", zap.Any("position", "init"))
}
