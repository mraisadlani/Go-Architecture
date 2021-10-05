package securities

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
	"time"
)

const AuthTokenValidTime = time.Hour * 3

func GenerateToken(username string) (string, error) {
	timer := time.Now().Add(AuthTokenValidTime).Unix()
	secret := viper.GetString("security.secret_key")
	claims := jwt.MapClaims{}

	claims["authorized"] = true
	claims["username"] = username
	claims["exp"] = timer

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(secret))

	if err != nil {
		return "", err
	}

	return tokenString, err
}