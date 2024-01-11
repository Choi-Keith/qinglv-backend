package jwtx_test

import (
	"qinglv-backend/pkg/jwtx"
	"testing"
)

func TestParseToken(t *testing.T) {
	secretKey := ""
	tokenString := ""
	jwtx.ParseToken(tokenString, secretKey)
}
