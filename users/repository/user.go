package repository

import (
	"Foundries/domain"
	"gorm.io/gorm"
)

type UserRepository struct {
	Conn *gorm.DB
}

// NewUserRepository will create an object that represent the interface
func NewUserRepository(Conn *gorm.DB) domain.UserRepository {
	return &UserRepository{Conn}
}

// Fetch will fetch all user data from database
func (u *UserRepository) Fetch() ([]domain.User, error) {
	var users []domain.User
	err := u.Conn.Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

// GetByID will get user data by given id
func (u *UserRepository) GetByID(id string) (domain.User, error) {
	var user domain.User
	err := u.Conn.Where("id = ?", id).First(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

// GetByEmail will get user data by given email
func (u *UserRepository) GetByEmail(email string) (domain.User, error) {
	var user domain.User
	err := u.Conn.Where("email = ?", email).First(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

// Store will store user data to database
func (u *UserRepository) Store(user *domain.User) error {
	err := u.Conn.Create(user).Error
	if err != nil {
		return err
	}
	return nil
}
