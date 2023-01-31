package services

import (
	"mlanches/src/models"
)

type UserRepository interface {
	FindAll() ([]models.User, error)
	FindOne(string) (models.User, error)
	Create(models.User) error
	Delete(string) (uint16, error)
	Update(string, models.User) (uint16, error)
}

type UserService struct {
	userRepository UserRepository
}

func NewUserService(userRepository UserRepository) *UserService {
	return &UserService{userRepository}
}

func (userService UserService) FindAll() ([]models.User, error) {
	result, err := userService.userRepository.FindAll()
	if err != nil {
		return []models.User{}, err
	}

	return result, nil
}

func (userService UserService) Create(user models.User) error {
	if err := userService.userRepository.Create(user); err != nil {
		return err
	}
	return nil
}

func (userService UserService) FindOne(id string) (models.User, error) {
	findUser, err := userService.userRepository.FindOne(id)
	if err != nil {
		return models.User{}, err
	}

	return findUser, nil
}

func (userService UserService) Delete(id string) (uint16, error) {
	result, err := userService.userRepository.Delete(id)
	if err != nil {
		return 0, err
	}
	return result, nil
}

func (userService UserService) Update(id string, userUpdate models.User) (uint16, error) {
	result, err := userService.userRepository.Update(id, userUpdate)
	if err != nil {
		return 0, err
	}

	return result, nil
}
