package server

import (
	"blog-service/pkg/config"
	"blog-service/pkg/db"
	"blog-service/pkg/logger"
	"blog-service/pkg/redis"
	"blog-service/pkg/routes"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func Run() {
	logger.InitLogger()
	defer logger.Logger.Sync()

	config.InitConfig()
	db.InitDB()
	redis.InitRedis()

	router := gin.Default()
	routes.InitRoutes(router)
	router.Run(viper.GetString("server.port"))
}
