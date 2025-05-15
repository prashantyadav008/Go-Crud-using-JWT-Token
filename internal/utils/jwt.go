package utils

import (
	"go-crud-msql/internal/models"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte("my-secret-key")

func GenerateJWT(userData *models.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{"userData": userData, "exp": time.Now().Add(time.Hour * 24).Unix()})

	return token.SignedString(jwtSecret)
}

func ValidateJWT(tokenString string) (map[string]interface{}, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	claims, ok := token.Claims.(jwt.MapClaims)

	if ok && token.Valid {

		// return whole userData object
		userData := claims["userData"].(map[string]interface{})

		// return whole userData object

		return userData, nil
	}

	return nil, err

}
