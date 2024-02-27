package jwtx

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/zeromicro/go-zero/core/logx"
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
		if v.Key == "userId" {
			claims["tokenId"] = strconv.FormatUint(v.Val.(uint64), 10)
		}
	}

	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}

func ParseToken(r *http.Request, secretKey string) (map[string]uint64, error) {
	if r.Header.Get("Authorization") == "" {
		return nil, errors.New("用户没有登录")
	}
	tokenString := strings.Split(r.Header.Get("Authorization"), " ")[1]
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secretKey), nil
	})
	if err != nil {
		return nil, err
	}
	m := make(map[string]uint64)
	var tokenId string
	for k, v := range token.Claims.(jwt.MapClaims) {
		if k == "userId" || k == "roleId" {
			m[k] = uint64(v.(float64))
		}
		if k == "tokenId" {
			tokenId = v.(string)
		}
	}
	newUserId, _ := strconv.ParseUint(tokenId, 10, 64)
	if err != nil {
		logx.Errorf("[ParseToken] failed: %+v\n", err)
	}
	m["userId"] = newUserId
	return m, nil
}
