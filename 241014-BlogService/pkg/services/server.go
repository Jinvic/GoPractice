package services

import (
	"blog-service/pkg/config"
	"blog-service/pkg/db"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func Run() {
	config.InitConfig()
	db.InitDB()

	router := gin.Default()
	router.Run(viper.GetString("server.port"))
}
