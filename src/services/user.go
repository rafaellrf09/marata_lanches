package services

import "mlanches/src/models"

type UserRepository interface {
	FindAll() ([]models.User, error)
	FindOne(string) (models.User, error)
	Create(models.User) error
	Delete(string) (int64, error)
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
	return models.User{}, nil
}
func (userService UserService) Delete(id string) (int64, error) {
	return 0, nil
}
