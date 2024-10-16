package user

import (
	"blog-service/pkg/define"
	"blog-service/pkg/logger"
	"blog-service/pkg/models"
	"blog-service/pkg/services/user"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// @Summary Register a new user
// @Description Register a new user with username, password, and email
// @Tags User
// @Accept json
// @Produce json
// @Param user body define.UserRegisterReq true "User registration details"
// @Success 200 {object} define.UserRegisterRes
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
func Register(c *gin.Context) {
	logger.Logger.Info("Register user")
	req := define.UserRegisterReq{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		logger.Logger.Error("Failed to bind JSON", zap.Any("position", "api"), zap.Error(err))
		return
	}

	// check if username already exists
	if ok, err := checkDuplicateUsername(req.Username); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		logger.Logger.Error("Failed to check duplicate username", zap.Any("position", "api"), zap.Error(err))
		return
	} else if ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username already exists"})
		logger.Logger.Error("Username already exists", zap.Any("position", "api"))
		return
	}

	// check if email already exists
	if ok, err := checkDuplicateEmail(req.Email); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		logger.Logger.Error("Failed to check duplicate email", zap.Any("position", "api"), zap.Error(err))
		return
	} else if ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email already exists"})
		logger.Logger.Error("Email already exists", zap.Any("position", "api"))
		return
	}

	u := models.User{
		Username: req.Username,
		Password: req.Password,
		Email:    req.Email,
	}

	// register user
	userInfo, err := user.Register(&u)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		logger.Logger.Error("Failed to register user", zap.Any("position", "api"), zap.Error(err))
		return
	}

	c.String(http.StatusOK, "User registered successfully")
	logger.Logger.Info("User registered successfully", zap.Any("position", "api"), zap.Any("user", userInfo))
}
