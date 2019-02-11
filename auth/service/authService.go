package service

import (
	"encoding/json"
	"github.com/edumar111/fastpv-auth/auth/core"
	"github.com/edumar111/fastpv-auth/auth/model"
	"github.com/edumar111/fastpv-auth/auth/parameters"
	"net/http"
)

//Login login with username and password and generate jwt
func Login(requestUser *model.UserLogin) (int, []byte) {
	authBackend :=core.InitJWTAuthenticationBackend()
	if authBackend.Authenticate(requestUser) {
		token, err := authBackend.GenerateToken(requestUser.Username)
		if err != nil {
			return http.StatusInternalServerError, []byte("")
		} else {
			response, _ := json.Marshal(parameters.TokenAuthentication{token})
			return http.StatusOK, response
		}
	}
	return http.StatusUnauthorized, []byte("")
}
