package user

import (
	"blog-service/pkg/logger"
	"blog-service/pkg/services/user"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func List(c *gin.Context) {
	userInfos, err := user.List()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		logger.Logger.Error("List users failed", zap.Any("error", err))
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": userInfos, "total": len(*userInfos)})
	logger.Logger.Info("List users successfully", zap.Any("total", len(*userInfos)))
}
