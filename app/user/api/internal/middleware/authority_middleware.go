package middleware

import (
	"errors"
	"fmt"
	"net/http"
	"qinglv-backend/app/user/api/internal/config"
	"qinglv-backend/app/user/rpc/client/userclass"
	"qinglv-backend/app/user/rpc/user"
	"qinglv-backend/common/globalKey"
	"qinglv-backend/common/response"
	"qinglv-backend/pkg/jwtx"
	"strings"
)

type AuthorityMiddleware struct {
	UserRpc userclass.UserClass
	c       *config.Config
}

func NewAuthorityMiddleware(userRpc userclass.UserClass, c *config.Config) *AuthorityMiddleware {
	return &AuthorityMiddleware{
		UserRpc: userRpc,
		c:       c,
	}
}

func (m *AuthorityMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO generate middleware implement function, delete after code implementation
		tokenMap, err := jwtx.ParseToken(r, m.c.JWTAuth.AccessSecret)
		if err != nil {
			response.FailCodeMsg(w, http.StatusUnauthorized, errors.New("请先登录"))
			return
		}
		tokenPrefix := fmt.Sprintf("%s%d", globalKey.TokenPrefixKey, tokenMap["userId"])
		tokenResp, err := m.UserRpc.GetToken(r.Context(), &user.GetTokenReq{
			TokenKey: tokenPrefix,
		})
		if err != nil {
			response.FailCodeMsg(w, http.StatusUnauthorized, errors.New("请先登录"))
			return
		}
		tokenString := strings.Split(r.Header.Get("Authorization"), " ")[1]
		if strings.Compare(tokenResp.Token, tokenString) != 0 {
			response.FailCodeMsg(w, http.StatusUnauthorized, errors.New("请重新登录"))
			return
		}
		// Passthrough to next handler if need
		next(w, r)

	}
}
