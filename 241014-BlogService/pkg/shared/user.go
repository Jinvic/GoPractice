package shared

import (
	"blog-service/pkg/define"
	"errors"

	"github.com/gin-gonic/gin"
)

func GetUserInfo(c *gin.Context) (*define.UserInfo, error) {
	userInfoAny, ok := c.Get("user_info")
	if !ok {
		return nil, errors.New("user info not found")
	}
	userInfo, ok := userInfoAny.(*define.UserInfo)
	if !ok {
		return nil, errors.New("user info not found")
	}
	return userInfo, nil
}
