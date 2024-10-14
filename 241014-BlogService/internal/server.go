package server

import (
	"blog-service/pkg/config"
	"blog-service/pkg/db"
	"blog-service/pkg/logger"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func Run() {
	logger.InitLogger()
	defer logger.Logger.Sync()

	config.InitConfig()
	db.InitDB()

	router := gin.Default()
	router.Run(viper.GetString("server.port"))
}
