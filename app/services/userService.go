package services

import (
	"go-delivery/app/models"
	"go-delivery/app/repositories"
)

type UserService struct {
}

var userRepository = new(repositories.UserRepository)

func (us *UserService) FindAll() ([]models.User, error) {
	return userRepository.FindAll()
}

func (us *UserService) FindById() {

}

func (us *UserService) FindByEmail() {

}

func (us *UserService) Create() {

}

func (us *UserService) Update() {

}

func (us *UserService) Delete() {

}
