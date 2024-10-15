package routes

import (
	"blog-service/pkg/db"
	"blog-service/pkg/define"
	"blog-service/pkg/logger"
	"blog-service/pkg/models"
	"blog-service/pkg/services/auth"
	"blog-service/pkg/services/user"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func registerUserRoutes(router *gin.Engine) {
	userGroup := router.Group("/users")
	{
		userGroup.POST("/register", register)
		userGroup.POST("/login", login)
		userGroup.POST("/logout", logout)
	}
}

func register(c *gin.Context) {
	logger.Logger.Info("Register user")
	req := define.UserRegisterReq{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		logger.Logger.Error("Failed to bind JSON", zap.Error(err))
		return
	}

	if ok, err := checkDuplicateUsername(req.Username); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		logger.Logger.Error("Failed to check duplicate username", zap.Error(err))
		return
	} else if ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username already exists"})
		logger.Logger.Error("Username already exists")
		return
	}
	if ok, err := checkDuplicateEmail(req.Email); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		logger.Logger.Error("Failed to check duplicate email", zap.Error(err))
		return
	} else if ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email already exists"})
		logger.Logger.Error("Email already exists")
		return
	}

	u := models.User{
		Username: req.Username,
		Password: req.Password,
		Email:    req.Email,
	}

	userInfo, err := user.Register(&u)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		logger.Logger.Error("Failed to register user", zap.Error(err))
		return
	}

	token, err := auth.GenerateToken(userInfo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		logger.Logger.Error("Failed to generate token", zap.Error(err))
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
	logger.Logger.Info("User registered successfully")
}

func login(c *gin.Context) {
}

func logout(c *gin.Context) {
}

func checkDuplicateUsername(username string) (bool, error) {
	var count int64
	err := db.DB.Model(&models.User{}).Where("username = ?", username).Count(&count).Error
	return count > 0, err
}

func checkDuplicateEmail(email string) (bool, error) {
	var count int64
	err := db.DB.Model(&models.User{}).Where("email = ?", email).Count(&count).Error
	return count > 0, err
}
