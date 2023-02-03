package services

import (
	"mlanches/src/models"
)

type AddressRepository interface {
	FindAll(string) ([]models.Address, error)
	FindOne(string) (models.Address, error)
	Create(models.Address) error
	Delete(string) (uint16, error)
	Update(string, models.Address) (uint16, error)
}

type AddressService struct {
	addressRepository AddressRepository
}

func NewAddressService(addressRepository AddressRepository) *AddressService {
	return &AddressService{addressRepository}
}

func (addressService AddressService) Create(address models.Address) error {
	// objectId, err := primitive.ObjectIDFromHex(string(address.UserId))
	// if err != nil {
	// 	return err
	// }

	// address.UserId = objectId

	if err := addressService.addressRepository.Create(address); err != nil {
		return err
	}
	return nil
}

func (addressService AddressService) FindOne(id string) (models.Address, error) {
	findUser, err := addressService.addressRepository.FindOne(id)
	if err != nil {
		return models.Address{}, err
	}

	return findUser, nil
}

func (addressService AddressService) FindAll(userId string) ([]models.Address, error) {
	result, err := addressService.addressRepository.FindAll(userId)
	if err != nil {
		return []models.Address{}, err
	}

	return result, nil
}

func (addresService AddressService) Delete(id string) (uint16, error) {
	return 0, nil
}

func (addresService AddressService) Update(id string, address models.Address) (uint16, error) {
	return 0, nil
}
