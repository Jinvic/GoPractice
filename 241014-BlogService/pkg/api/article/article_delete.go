package article

import (
	"blog-service/pkg/logger"
	"blog-service/pkg/services/article"
	"blog-service/pkg/shared"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Delete(c *gin.Context) {
	logger.Logger.Info("Delete article")
	articleID, err := shared.GetArticleID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		logger.Logger.Error("invalid article id", zap.Any("position", "api"), zap.Error(err))
		return
	}

	err = article.Delete(articleID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		logger.Logger.Error("failed to delete article", zap.Any("position", "api"), zap.Error(err))
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "article deleted"})
	logger.Logger.Info("article deleted", zap.Any("article_id", articleID))
}
