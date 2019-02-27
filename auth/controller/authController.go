package controller

import (
	"encoding/json"
	"github.com/edumar111/fastpv-auth/auth/model"
	"github.com/edumar111/fastpv-auth/auth/service"
	"log"
	"net/http"
)

//Login login User
func Login(w http.ResponseWriter , r *http.Request){
	log.Println("init Login - controller")
	requestUser := new(model.UserLogin)
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&requestUser)

	responseStatus, token := service.Login(requestUser)
		w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(responseStatus)
	w.Write(token)
}
func RefreshToken(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	requestUser := new(model.UserLogin)
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&requestUser)

	w.Header().Set("Content-Type", "application/json")
	w.Write(service.RefreshToken(requestUser))
}
func Logout(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	err := service.Logout(r)
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusOK)
	}
}