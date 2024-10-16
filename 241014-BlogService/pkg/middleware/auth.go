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
			c.JSON(http.StatusUnauthorized, gin.H{"error": "token not provided"})
			logger.Logger.Error("token not provided", zap.Any("position", "middleware"))
			c.Abort()
			return
		}

		if isBanned, err := auth.IsBanned(tokenString); err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "failed to check token"})
			logger.Logger.Error("failed to check token", zap.Any("position", "middleware"), zap.Error(err))
			c.Abort()
			return
		} else if isBanned {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "token is banned"})
			logger.Logger.Error("token is banned", zap.Any("position", "middleware"))
			c.Abort()
			return
		}

		claims, err := auth.ParseToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			logger.Logger.Error("invalid token", zap.Any("position", "middleware"), zap.Error(err))
			c.Abort()
			return
		}

		c.Set("user_info", claims.UserInfo)
		c.Next()
	}
}
