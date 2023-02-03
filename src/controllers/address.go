package controllers

import (
	"encoding/json"
	"io/ioutil"
	"mlanches/src/models"
	"mlanches/src/responses"
	"net/http"
)

type AddressService interface {
	Create(models.Address) error
	FindOne(string) (models.Address, error)
	FindAll(string) ([]models.Address, error)
	Delete(string) (uint16, error)
	Update(string, models.Address) (uint16, error)
}

type AddressController struct {
	addressService AddressService
}

func NewAddressController(addressService AddressService) *AddressController {
	return &AddressController{addressService}
}

func (addressController AddressController) Create(w http.ResponseWriter, r *http.Request) {
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	defer r.Body.Close()

	var newAddress models.Address
	if err = json.Unmarshal(requestBody, &newAddress); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	if err = addressController.addressService.Create(newAddress); err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusCreated, nil)
}

func (addressController AddressController) FindAll(w http.ResponseWriter, r *http.Request) {
	userId := r.URL.Query().Get("userId") //deve virar um parametro vindo do jwt dentro da requisicao
	result, err := addressController.addressService.FindAll(userId)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
	}
	responses.JSON(w, http.StatusOK, result)
}

func (addressController AddressController) Delete(w http.ResponseWriter, r *http.Request) {

}
