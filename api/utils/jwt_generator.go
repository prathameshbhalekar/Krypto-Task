package utils

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"github.com/spf13/viper"
)

func GenerateNewAccessToken() (string, error) {
	secret := viper.GetString("JWT_SECRET_KEY")

	minutesCount := viper.GetInt("JWT_SECRET_KEY_EXPIRE_MINUTES_COUNT")

	claims := jwt.MapClaims{
		"user_id": uuid.New().String(),
	}

	claims["exp"] = time.Now().Add(time.Minute * time.Duration(minutesCount)).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return t, nil
}
