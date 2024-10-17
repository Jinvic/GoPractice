package article

import (
	"blog-service/pkg/define"
	"blog-service/pkg/logger"
	"blog-service/pkg/services/article"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Edit(c *gin.Context) {
	logger.Logger.Info("Edit article")
	req := define.ArticleEditReq{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		logger.Logger.Error("invalid article edit request", zap.Any("position", "api"), zap.Error(err))
		return
	}

	err = article.Edit(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		logger.Logger.Error("failed to edit article", zap.Any("position", "api"), zap.Error(err))
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "article edited"})
	logger.Logger.Info("article edited", zap.Any("article_id", req.ID))
}
