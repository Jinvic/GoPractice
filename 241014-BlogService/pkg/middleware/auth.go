package middleware

import (
	"blog-service/pkg/logger"
	"blog-service/pkg/services/auth"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "未提供令牌"})
			logger.Logger.Error("未提供令牌", zap.Any("position", "middleware"))
			c.Abort()
			return
		}

		claims, err := auth.ParseToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "无效的令牌"})
			logger.Logger.Error("无效的令牌", zap.Any("position", "middleware"), zap.Error(err))
			c.Abort()
			return
		}

		c.Set("user_info", claims.UserInfo)
		c.Next()
	}
}
