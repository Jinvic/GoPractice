package article

import (
	"blog-service/pkg/logger"
	"blog-service/pkg/services/article"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func ListAll(c *gin.Context) {
	logger.Logger.Info("List all articles")
	articles, err := article.ListAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		logger.Logger.Error("failed to list all articles", zap.Any("position", "api"), zap.Error(err))
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": articles, "total": len(*articles)})
	logger.Logger.Info("articles listed", zap.Any("total", len(*articles)))
}
