package model

import "net/http"

type Route struct {
	Name       string
	Method     string
	Pattern    string
	HandleFunc http.HandlerFunc
}
type RouteHandle struct {
	Name       string
	Method     string
	Pattern    string
	Handle http.Handler
}

type Routes []Route
type RoutesHandle []RouteHandle