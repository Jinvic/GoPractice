package article

import (
	"blog-service/pkg/define"
	"blog-service/pkg/logger"
	"blog-service/pkg/services/article"
	"blog-service/pkg/shared"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Create(c *gin.Context) {
	logger.Logger.Info("Create article")
	req := define.ArticleCreateReq{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		logger.Logger.Error("invalid article create request", zap.Any("position", "api"), zap.Error(err))
		return
	}

	userInfo, err := shared.GetUserInfo(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		logger.Logger.Error("user info not found", zap.Any("position", "api"), zap.Error(err))
		return
	}

	articleID, err := article.Create(userInfo.ID, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		logger.Logger.Error("failed to create article", zap.Any("position", "api"), zap.Error(err))
		return
	}

	c.JSON(http.StatusOK, gin.H{"article_id": articleID})
	logger.Logger.Info("article created", zap.Any("article_id", articleID))
}
