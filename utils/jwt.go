package utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwt_secret string = os.Getenv("JWT_SECRET")

func CreateToken(user_id int, is_admin bool) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":  user_id,
		"is_admin": is_admin,
		"exp":      time.Now().Add(time.Hour * 2).Unix(),
	})

	// Sign the token
	tokenString, err := token.SignedString([]byte(jwt_secret))

	if err != nil {
		return "", err
	}

	return tokenString, nil

}

func DecodeToken(jwt_token string) (int, bool, error) {
	token, err := jwt.Parse(jwt_token, func(token *jwt.Token) (interface{}, error) {

		return []byte(jwt_secret), nil
	})
	if err != nil {
		return 0, false, err
	}

	claims := token.Claims.(jwt.MapClaims)
	user_id := int(claims["user_id"].(float64))
	is_admin := claims["is_admin"].(bool)
	exp := int64(claims["exp"].(float64))
	if time.Now().Unix() > exp {
		return 0, false, jwt.ErrTokenExpired
	}

	return user_id, is_admin, nil
}
