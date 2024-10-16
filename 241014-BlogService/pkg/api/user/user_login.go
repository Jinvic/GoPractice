package user

import (
	"blog-service/pkg/define"
	"blog-service/pkg/logger"
	"blog-service/pkg/services/auth"
	"blog-service/pkg/services/user"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// @Summary Login user
// @Description Login user with username and password
// @Tags User
// @Accept json
// @Produce json
// @Param user body define.UserLoginReq true "User login details"
// @Success 200 {object} define.UserLoginRes
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
func Login(c *gin.Context) {
	logger.Logger.Info("Login user")
	req := define.UserLoginReq{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		logger.Logger.Error("Failed to bind JSON", zap.Error(err))
		return
	}

	// login user
	userInfo, err := user.Login(req.Username, req.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		logger.Logger.Error("Failed to login user", zap.Error(err))
		return
	}
	if userInfo == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		logger.Logger.Error("Invalid username or password")
		return
	}

	// check if token already exists
	if hasToken, err := auth.HasToken(userInfo.ID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		logger.Logger.Error("Failed to check if token exists", zap.Error(err))
		return
	} else if hasToken {
		// ban old token
		oldToken, err := auth.GetToken(userInfo.ID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			logger.Logger.Error("Failed to get old token", zap.Error(err))
			return
		}

		claims, err := auth.ParseToken(oldToken)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			logger.Logger.Error("Failed to parse old token", zap.Error(err))
			return
		}

		err = auth.BanToken(oldToken, claims.RegisteredClaims.ExpiresAt.Time)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			logger.Logger.Error("Failed to ban old token", zap.Error(err))
			return
		}
	}

	// generate new token
	expiredAt := getExpiredAt()
	token, err := auth.GenerateToken(userInfo, expiredAt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		logger.Logger.Error("Failed to generate token", zap.Error(err))
		return
	}

	// set new token
	err = auth.SetToken(token, userInfo.ID, expiredAt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		logger.Logger.Error("Failed to set new token", zap.Error(err))
		return
	}

	c.JSON(http.StatusOK, gin.H{"user_info": userInfo, "token": token})
	logger.Logger.Info("User logged in successfully", zap.Any("user_info", userInfo))
}
