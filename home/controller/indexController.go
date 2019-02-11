package controller


import (
	"fmt"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from server web Go")
}

func  Test(w http.ResponseWriter, r *http.Request,next http.HandlerFunc){
	fmt.Fprintf(w, "Test ok")
}