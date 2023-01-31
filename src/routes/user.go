package routes

import (
	"mlanches/src/controllers"
	"net/http"

	"github.com/gorilla/mux"
)

func InsertUserRoutes(router *mux.Router, UserController *controllers.UserController) {
	UserRoutes := []Route{
		{
			URL:    "/users",
			Method: http.MethodPost,
			Func:   UserController.Create,
			isAuth: false,
		},
		{
			URL:    "/users",
			Method: http.MethodGet,
			Func:   UserController.FindAll,
			isAuth: false,
		},
		{
			URL:    "/users/{id}",
			Method: http.MethodGet,
			Func:   UserController.FindOne,
			isAuth: false,
		},
		{
			URL:    "/users/{id}",
			Method: http.MethodDelete,
			Func:   UserController.Delete,
			isAuth: false,
		},
		{
			URL:    "/users/{id}",
			Method: http.MethodPut,
			Func:   UserController.Update,
			isAuth: false,
		},
	}

	for _, route := range UserRoutes {
		router.HandleFunc(route.URL, route.Func).Methods(route.Method)
	}
}
