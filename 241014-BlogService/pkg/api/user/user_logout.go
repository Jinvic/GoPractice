package user

import (
	"blog-service/pkg/logger"
	"blog-service/pkg/shared"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Logout(c *gin.Context) {
	logger.Logger.Info("Logout user")

	userInfo, err := shared.GetUserInfo(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user info not found"})
		logger.Logger.Error("user info not found", zap.Any("position", "api"), zap.Error(err))
		return
	}

	err = banOldToken(userInfo.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to logout"})
		logger.Logger.Error("Failed to logout", zap.Any("position", "api"), zap.Error(err))
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Logout successfully"})
	logger.Logger.Info("Logout successfully", zap.Any("user_info", userInfo))
}
