package routes

import (
	"mlanches/src/controllers"
	"net/http"

	"github.com/gorilla/mux"
)

func InsertAddressRoutes(router *mux.Router, AddressController *controllers.AddressController) {
	AddressRoutes := []Route{
		{
			URL:    "/address",
			Method: http.MethodPost,
			Func:   AddressController.Create,
			isAuth: false,
		},
		{
			URL:    "/address",
			Method: http.MethodGet,
			Func:   AddressController.FindAll,
			isAuth: false,
		},
	}

	for _, route := range AddressRoutes {
		router.HandleFunc(route.URL, route.Func).Methods(route.Method)
	}
}
