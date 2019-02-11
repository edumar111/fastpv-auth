package routes

import (
	"github.com/edumar111/fastpv-auth/auth/core"

	"github.com/urfave/negroni"
)

func FilterMiddleware(funparm negroni.HandlerFunc ) *negroni.Negroni {
	router := negroni.New(
		negroni.HandlerFunc(core.RequireTokenAuthentication),
		negroni.HandlerFunc(funparm),
	)
	return router
}
