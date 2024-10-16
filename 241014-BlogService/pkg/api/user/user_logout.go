package user

import (
	"blog-service/pkg/define"
	"blog-service/pkg/logger"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Logout(c *gin.Context) {
	logger.Logger.Info("Logout user")

	userInfo, ok := c.Get("user_info")
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "token not provided"})
		logger.Logger.Error("token not provided", zap.Any("position", "api"))
		return
	}
	err := banOldToken(userInfo.(define.UserInfo).ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to logout"})
		logger.Logger.Error("Failed to logout", zap.Any("position", "api"), zap.Error(err))
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Logout successfully"})
	logger.Logger.Info("Logout successfully", zap.Any("user_info", userInfo))
}
