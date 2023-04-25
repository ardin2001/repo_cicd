package repositories

import (
	"unit_testing/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	GetUsersRepository() ([]*models.User, error)
	GetUserRepository(id string) (*models.User, error)
	CreateRepository(user models.User) (*models.User, error)
	UpdateRepository(id string, userBody models.User) (*models.User, error)
	DeleteRepository(id string) error
}

type UserRepo struct {
	DB *gorm.DB
}

func NewUserRepository(DB *gorm.DB) UserRepository {
	return &UserRepo{
		DB: DB,
	}
}

func (u *UserRepo) GetUsersRepository() ([]*models.User, error) {
	var users []*models.User

	if err := u.DB.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func (u *UserRepo) GetUserRepository(id string) (*models.User, error) {
	var user models.User

	if err := u.DB.Where("ID = ?", id).Take(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *UserRepo) CreateRepository(user models.User) (*models.User, error) {
	if err := u.DB.Save(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *UserRepo) UpdateRepository(id string, userBody models.User) (*models.User, error) {
	user, err := u.GetUserRepository(id)
	if err != nil {
		return nil, err
	}

	err = u.DB.Where("ID = ?", id).Updates(models.User{Name: userBody.Name, Email: userBody.Email, Password: userBody.Password}).Error
	if err != nil {
		return nil, err
	}

	user.Name = userBody.Name
	user.Email = userBody.Email
	user.Password = userBody.Password

	return user, nil
}

func (u *UserRepo) DeleteRepository(id string) error {
	_, err := u.GetUserRepository(id)
	if err != nil {
		return err
	}

	if err := u.DB.Delete(&models.User{}, id).Error; err != nil {
		return err
	}

	return nil
}
