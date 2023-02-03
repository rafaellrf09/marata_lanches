package services

import (
	"fmt"
	"mlanches/src/models"
	"mlanches/src/services"
	"strings"
	"testing"
)

type userRepositoryMock struct{}

func (repo userRepositoryMock) FindAll() ([]models.User, error) {
	return []models.User{
		{
			Email:    "rafaellrf09@gmail.com",
			Password: "123456",
			Name:     "Rafael Figueiredo",
			Phone:    "22981393794",
		}, {
			Email:    "wagao@gmail.com",
			Password: "123456",
			Name:     "wagao",
			Phone:    "22222222222",
		},
	}, nil
}
func (repo userRepositoryMock) FindOne(id string) (models.User, error) {
	if id == "11111" {
		return models.User{}, fmt.Errorf("user not found")
	}

	return models.User{
		Email:    "rafaellrf09@gmail.com",
		Password: "123456",
		Name:     "Rafael Figueiredo",
		Phone:    "22981393794",
	}, nil
}
func (repo userRepositoryMock) Create(user models.User) error {
	if strings.Contains(user.Name, "wagao") {
		return fmt.Errorf("Error to create user")
	}
	return nil
}
func (repo userRepositoryMock) Delete(id string) (uint16, error) {
	if id == "11111" {
		return 0, fmt.Errorf("user not found")
	}
	return 1, nil
}
func (repo userRepositoryMock) Update(id string, userUpdate models.User) (uint16, error) {
	if id == "11111" {
		return 0, fmt.Errorf("user not found")
	}
	return 1, nil
}

func TestCreate(t *testing.T) {
	userRepository := userRepositoryMock{}
	userService := services.NewUserService(userRepository)

	userSuccess := models.User{
		Email:    "rafaellrf09@gmail.com",
		Password: "123456",
		Name:     "Rafael Figueiredo",
		Phone:    "22981393794",
	}

	userError := models.User{
		Email:    "wagao@gmail.com",
		Password: "123456",
		Name:     "wagao",
		Phone:    "22222222222",
	}

	if err := userService.Create(userSuccess); err != nil {
		t.Error(err)
	}

	if err := userService.Create(userError); err == nil {
		t.Error(err)
	}
}

func TestFindAll(t *testing.T) {
	userRepository := userRepositoryMock{}
	userService := services.NewUserService(userRepository)

	_, err := userService.FindAll()
	if err != nil {
		t.Error(err)
	}
}

func TestFindOne(t *testing.T) {
	userRepository := userRepositoryMock{}
	userService := services.NewUserService(userRepository)

	_, err := userService.FindOne("12345")
	if err != nil {
		t.Error(err)
	}

	_, err = userService.FindOne("11111")
	if err == nil {
		t.Error(err)
	}
}

func TestDelete(t *testing.T) {
	userRepository := userRepositoryMock{}
	userService := services.NewUserService(userRepository)

	_, err := userService.Delete("12345")
	if err != nil {
		t.Error(err)
	}

	_, err = userService.Delete("11111")
	if err == nil {
		t.Error(err)
	}
}

func TestUpdate(t *testing.T) {
	userRepository := userRepositoryMock{}
	userService := services.NewUserService(userRepository)

	userSuccess := models.User{
		Email:    "rafaellrf09@gmail.com",
		Password: "123456",
		Name:     "Rafael Figueiredo",
		Phone:    "22981393794",
	}

	userError := models.User{
		Email:    "wagao@gmail.com",
		Password: "123456",
		Name:     "wagao",
		Phone:    "22222222222",
	}

	_, err := userService.Update("12345", userSuccess)
	if err != nil {
		t.Error(err)
	}

	_, err = userService.Update("11111", userError)
	if err == nil {
		t.Error(err)
	}
}
