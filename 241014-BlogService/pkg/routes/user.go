package routes

import (
	"blog-service/pkg/db"
	"blog-service/pkg/define"
	"blog-service/pkg/logger"
	"blog-service/pkg/models"
	"blog-service/pkg/services/auth"
	"blog-service/pkg/services/user"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
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

// @Summary Register a new user
// @Description Register a new user with username, password, and email
// @Tags User
// @Accept json
// @Produce json
// @Param user body define.UserRegisterReq true "User registration details"
// @Success 200 {object} define.UserRegisterRes
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
func register(c *gin.Context) {
	logger.Logger.Info("Register user")
	req := define.UserRegisterReq{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		logger.Logger.Error("Failed to bind JSON", zap.Error(err))
		return
	}

	// check if username already exists
	if ok, err := checkDuplicateUsername(req.Username); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		logger.Logger.Error("Failed to check duplicate username", zap.Error(err))
		return
	} else if ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username already exists"})
		logger.Logger.Error("Username already exists")
		return
	}

	// check if email already exists
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

	// register user
	userInfo, err := user.Register(&u)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		logger.Logger.Error("Failed to register user", zap.Error(err))
		return
	}

	c.String(http.StatusOK, "User registered successfully")
	logger.Logger.Info("User registered successfully", zap.Any("user", userInfo))
}

// @Summary Login user
// @Description Login user with username and password
// @Tags User
// @Accept json
// @Produce json
// @Param user body define.UserLoginReq true "User login details"
// @Success 200 {object} define.UserLoginRes
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
func login(c *gin.Context) {
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

func getExpiredAt() time.Time {
	expiredHour := viper.GetInt("token.expired.hour")
	expiredMinute := viper.GetInt("token.expired.minute")
	expiredSecond := viper.GetInt("token.expired.second")
	return time.Now().Add(time.Hour*time.Duration(expiredHour) +
		time.Minute*time.Duration(expiredMinute) +
		time.Second*time.Duration(expiredSecond))
}
