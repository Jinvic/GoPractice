package auth

import (
	"blog-service/pkg/define"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/spf13/viper"
)

func GenerateToken(userInfo *define.UserInfo) (string, error) {
	claims := define.UserClaims{
		UserInfo: userInfo,
		RegisteredClaims: &jwt.RegisteredClaims{
			Issuer:    "blog-service",
			Subject:   "user_token",
			Audience:  []string{userInfo.Username},
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
			NotBefore: jwt.NewNumericDate(time.Now()),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ID:        uuid.NewString(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secret := viper.GetString("jwt.secret")
	return token.SignedString([]byte(secret))
}

func ParseToken(tokenString string) (*define.UserClaims, error) {
	secret := viper.GetString("jwt.secret")
	token, err := jwt.ParseWithClaims(tokenString, &define.UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}
	return token.Claims.(*define.UserClaims), nil
}
