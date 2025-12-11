package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secret = []byte("SECRET_SUPER_KEY_1234567890@#$^&*")

func GenerateToken(userId int) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userId,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,claims)

	validToken,err := token.SignedString(secret)

	if err !=nil{
		return "", err
	}
	return validToken, nil
}


func ParseToken(tokenStr string)(jwt.MapClaims,error){
	token,err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return secret,nil
	})
	if err !=nil || !token.Valid{
		return nil,err
	}

	return token.Claims.(jwt.MapClaims),nil
}
