package main

import (
	"fmt"
	"log"
	"mlanches/src/config"
	"mlanches/src/controllers"
	"mlanches/src/database"
	"mlanches/src/respository"
	"mlanches/src/routes"
	"mlanches/src/services"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	config.Load()

	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}

	defer database.Close(db)

	router := mux.NewRouter()

	userCollection := database.DbConnect(db, "users")
	userRepository := respository.NewUserRepository(userCollection)
	userService := services.NewUserService(userRepository)
	userController := controllers.NewUserController(userService)
	routes.InsertUserRoutes(router, userController)

	addressCollection := database.DbConnect(db, "addresses")
	addressRepository := respository.NewAddessRepository(addressCollection)
	addressService := services.NewAddressService(addressRepository)
	addressController := controllers.NewAddressController(addressService)

	routes.InsertAddressRoutes(router, addressController)

	fmt.Println("Sever Listen Port 5000")
	http.ListenAndServe("localhost:5000", router)
}
