package controller

import (
	"encoding/json"
	"fmt"
	"github.com/edumar111/fastpv-auth/user/model"
	"github.com/edumar111/fastpv-auth/user/service"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

// CreateUser crear usuario
func CreateUser(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	user := model.User{}
	err := decoder.Decode(&user)
	if err != nil{
		w.WriteHeader(404)
		return
	}

	fmt.Println("CreateUser", user)

	userUpdate,err := service.CreateUser(user)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	json.NewEncoder(w).Encode(userUpdate)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	iduser, err := strconv.ParseInt(params["id"], 10, 64)

	if err != nil{
		w.WriteHeader(404)
		return
	}
	user :=service.GetUser(iduser)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(user)
}