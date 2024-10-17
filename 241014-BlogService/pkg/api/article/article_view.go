package article

import (
	"blog-service/pkg/logger"
	"blog-service/pkg/services/article"
	"blog-service/pkg/shared"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func View(c *gin.Context) {
	logger.Logger.Info("View article")
	articleID, err := shared.GetArticleID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		logger.Logger.Error("invalid article id", zap.Any("position", "api"), zap.Error(err))
		return
	}

	articleInfo, err := article.View(articleID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		logger.Logger.Error("failed to view article", zap.Any("position", "api"), zap.Error(err))
		return
	}

	c.JSON(http.StatusOK, articleInfo)
	logger.Logger.Info("article viewed", zap.Any("article_id", articleID))
}
