package auth

import (
	"blog-service/pkg/define"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/spf13/viper"
)

func GenerateToken(userInfo *define.UserInfo) (string, error) {
	expiredAt := time.Now().Add(time.Hour * 24)
	claims := define.UserClaims{
		UserInfo: userInfo,
		RegisteredClaims: &jwt.RegisteredClaims{
			Issuer:    "blog-service",
			Subject:   "user_token",
			Audience:  []string{userInfo.Username},
			ExpiresAt: jwt.NewNumericDate(expiredAt),
			NotBefore: jwt.NewNumericDate(time.Now()),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ID:        uuid.NewString(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secret := viper.GetString("jwt.secret")
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	err = SetToken(tokenString, userInfo.ID)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ParseToken(tokenString string) (*define.UserClaims, error) {
	secret := viper.GetString("jwt.secret")
	token, err := jwt.ParseWithClaims(tokenString, &define.UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*define.UserClaims)
	if !ok {
		return nil, errors.New("invalid token claims")
	}

	banned, err := IsBannedToken(tokenString)
	if err != nil {
		return nil, err
	}
	if banned {
		return nil, errors.New("token banned")
	}

	return claims, nil
}