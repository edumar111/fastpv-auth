package service

import (
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
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
			fmt.Println("Ok")
			return http.StatusOK, response
		}
	}
	return http.StatusUnauthorized, []byte("")
}
func RefreshToken(requestUser *model.UserLogin) []byte {
	authBackend := core.InitJWTAuthenticationBackend()
	token, err := authBackend.GenerateToken(requestUser.Username)
	if err != nil {
		panic(err)
	}
	response, err := json.Marshal(parameters.TokenAuthentication{token})
	if err != nil {
		panic(err)
	}
	return response
}

func Logout(req *http.Request) error {
	authBackend := core.InitJWTAuthenticationBackend()
	tokenRequest, err := request.ParseFromRequest(req, request.OAuth2Extractor, func(token *jwt.Token) (interface{}, error) {
		return authBackend.PublicKey, nil
	})
	if err != nil {
		return err
	}
	tokenString := req.Header.Get("Authorization")
	err = authBackend.Logout(tokenString, tokenRequest)
	if err != nil {
		fmt.Println(err)
	}
	return err
}
