package routes

import (
	"mlanches/src/controllers"

	"github.com/gorilla/mux"
)

func Generate(r *mux.Router, userController *controllers.UserController) *mux.Router {
	InsertUserRoutes(r, userController)
	return r
}
