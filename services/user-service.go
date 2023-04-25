package services

import (
	"unit_testing/models"
	"unit_testing/repositories"
)

type UserService interface {
	GetUsersService() ([]*models.User, error)
	GetUserService(id string) (*models.User, error)
	CreateService(user models.User) (*models.User, error)
	UpdateService(id string, userBody models.User) (*models.User, error)
	DeleteService(id string) error
}

type Users struct {
	userR repositories.UserRepository
}

func NewUserService(userR repositories.UserRepository) UserService {
	return &Users{
		userR: userR,
	}
}

func (u *Users) GetUsersService() ([]*models.User, error) {
	users, err := u.userR.GetUsersRepository()
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (u *Users) GetUserService(id string) (*models.User, error) {
	user, err := u.userR.GetUserRepository(id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *Users) CreateService(user models.User) (*models.User, error) {
	userR, err := u.userR.CreateRepository(user)
	if err != nil {
		return nil, err
	}

	return userR, nil
}

func (u *Users) UpdateService(id string, userBody models.User) (*models.User, error) {
	user, err := u.userR.UpdateRepository(id, userBody)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *Users) DeleteService(id string) error {
	err := u.userR.DeleteRepository(id)
	if err != nil {
		return err
	}

	return nil
}
