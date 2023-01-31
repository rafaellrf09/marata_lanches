package controllers

import (
	"encoding/json"
	"io/ioutil"
	"mlanches/src/models"
	"mlanches/src/responses"
	"net/http"

	"github.com/gorilla/mux"
)

type UserService interface {
	Create(models.User) error
	FindOne(string) (models.User, error)
	FindAll() ([]models.User, error)
	Delete(string) (uint16, error)
	Update(string, models.User) (uint16, error)
}

type UserController struct {
	userService UserService
}

func NewUserController(userService UserService) *UserController {
	return &UserController{userService}
}

func (userController UserController) Create(w http.ResponseWriter, r *http.Request) {
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	defer r.Body.Close()

	var newUser models.User
	if err = json.Unmarshal(requestBody, &newUser); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	if err = userController.userService.Create(newUser); err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusCreated, nil)
}

func (userController UserController) FindAll(w http.ResponseWriter, r *http.Request) {
	result, err := userController.userService.FindAll()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
	}
	responses.JSON(w, http.StatusOK, result)
}

func (userController UserController) FindOne(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	result, err := userController.userService.FindOne(id)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
	}
	responses.JSON(w, http.StatusOK, result)
}

func (userController UserController) Delete(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	result, err := userController.userService.Delete(id)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
	}
	responses.JSON(w, http.StatusOK, result)
}

func (userController UserController) Update(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	defer r.Body.Close()

	var updateUser models.User
	if err = json.Unmarshal(requestBody, &updateUser); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	result, err := userController.userService.Update(id, updateUser)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
	}
	responses.JSON(w, http.StatusOK, result)
}
