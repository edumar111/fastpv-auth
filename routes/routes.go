package routes

import (
	authController "github.com/edumar111/fastpv-auth/auth/controller"
	homeController "github.com/edumar111/fastpv-auth/home/controller"
	"github.com/edumar111/fastpv-auth/routes/model"

	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router{
	router :=mux.NewRouter().StrictSlash(true)
	for _, route := range routes{
		router.Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandleFunc)
	}
	return  router
}

var routes = model.Routes{
	model.Route{
		"index",
		"GET",
		"/",
		homeController.Index,
	},
	model.Route{
		"Login",
		"POST",
		"/login",
		authController.Login,
	},
}