package user

import (
	"blog-service/pkg/define"
	"blog-service/pkg/logger"
	"blog-service/pkg/services/user"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Delete(c *gin.Context) {
	logger.Logger.Info("Delete user")
	userInfoAny, ok := c.Get("user_info")
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "user info not found"})
		logger.Logger.Error("user info not found", zap.Any("position", "api"))
		return
	}
	userInfo, ok := userInfoAny.(*define.UserInfo)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "user info not found"})
		logger.Logger.Error("user info not found", zap.Any("position", "api"))
		return
	}

	err := banOldToken(userInfo.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		logger.Logger.Error("Failed to ban old token", zap.Any("position", "api"), zap.Error(err))
		return
	}

	err = user.Delete(userInfo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		logger.Logger.Error("Failed to delete user", zap.Any("position", "api"), zap.Error(err))
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
	logger.Logger.Info("User deleted successfully", zap.Any("user_info", userInfo))
}
