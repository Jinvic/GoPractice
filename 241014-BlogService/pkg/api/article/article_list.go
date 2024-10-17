package article

import (
	"blog-service/pkg/logger"
	"blog-service/pkg/services/article"
	"blog-service/pkg/shared"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func List(c *gin.Context) {
	logger.Logger.Info("List articles")
	userInfo, err := shared.GetUserInfo(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		logger.Logger.Error("invalid user info", zap.Any("position", "api"), zap.Error(err))
		return
	}

	articles, err := article.List(userInfo.ID, userInfo.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		logger.Logger.Error("failed to list articles", zap.Any("position", "api"), zap.Error(err))
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": articles, "total": len(*articles)})
	logger.Logger.Info("articles listed", zap.Any("user_id", userInfo.ID))
}
