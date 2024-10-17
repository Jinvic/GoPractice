package middleware

import (
	"blog-service/pkg/logger"
	"blog-service/pkg/services/auth"
	"blog-service/pkg/shared"
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

func AdminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		userInfo, err := shared.GetUserInfo(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "user info not found"})
			logger.Logger.Error("user info not found", zap.Any("position", "middleware"), zap.Error(err))
			c.Abort()
			return
		}

		if userInfo.IsAdmin() {
			c.Next()
		} else {
			c.JSON(http.StatusForbidden, gin.H{"error": "permission denied"})
			logger.Logger.Error("permission denied", zap.Any("position", "middleware"))
			c.Abort()
		}
	}
}

func OwnershipMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
	}
}
