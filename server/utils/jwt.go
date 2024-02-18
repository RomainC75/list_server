package utils

import (
	"fmt"

	"github.com/RomainC75/todo2/config"
	"github.com/golang-jwt/jwt"
)

type UserClaims struct {
	Id    uint   `json:"id"`
	Email string `json:"email"`
	jwt.StandardClaims
}

func NewAccessToken(claims UserClaims) (string, error) {
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return accessToken.SignedString([]byte(config.Get().Jwt.Secret))
}

func NewRefreshToken(claims jwt.StandardClaims) (string, error) {
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return refreshToken.SignedString([]byte(config.Get().Jwt.Secret))
}

func ParseAccessToken(accessToken string) (*jwt.Claims, error) {
	parsedAccessToken, err := jwt.ParseWithClaims(accessToken, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.Get().Jwt.Secret), nil
	})
	if err != nil {
		return nil, err
	}
	return &parsedAccessToken.Claims, nil
}

func GetClaimsFromToken(tokenString string) (jwt.MapClaims, error) {
	secret := config.Get().Jwt.Secret
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, err
}

func ParseRefreshToken(refreshToken string) *jwt.StandardClaims {
	parsedRefreshToken, _ := jwt.ParseWithClaims(refreshToken, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.Get().Jwt.Secret), nil
	})
	return parsedRefreshToken.Claims.(*jwt.StandardClaims)
}
