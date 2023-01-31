package routes

import "net/http"

type Route struct {
	URL    string
	Method string
	Func   func(w http.ResponseWriter, r *http.Request)
	isAuth bool
}
