package jwtx

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// Option describes the jwt extra data
type Option struct {
	Key string
	Val any
}

// WithOption returns the option from k/v
func WithOption(key string, val any) Option {
	return Option{
		Key: key,
		Val: val,
	}
}

// NewJwtToken returns the jwt token from the given data.
func NewJwtToken(secretKey string, seconds int64, opt ...Option) (string, error) {
	claims := make(jwt.MapClaims)
	iat := time.Now().Unix()
	claims["exp"] = iat + seconds
	claims["iat"] = iat

	for _, v := range opt {
		claims[v.Key] = v.Val
	}

	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}

func ParseToken(tokenString, secretKey string) ([]Option, error) {
	token, err := jwt.ParseWithClaims(tokenString, &jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	if err != nil {
		return nil, err
	}
	fmt.Printf("token claims: %+v\n", token.Claims)
	return nil, nil
}
