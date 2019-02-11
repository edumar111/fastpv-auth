package controller

import (
	"encoding/json"
	"fmt"
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
	fmt.Println(requestUser)
	responseStatus, token := service.Login(requestUser)
		w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(responseStatus)

	w.Write(token)
}