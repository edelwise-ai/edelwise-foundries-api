package repository

import (
	"Foundries/domain"
	"github.com/gin-gonic/gin"
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
func (u *UserRepository) Fetch(c *gin.Context) ([]domain.User, error) {
	var users []domain.User
	err := u.Conn.Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

// GetByID will get user data by given id
func (u *UserRepository) GetByID(c *gin.Context, id string) (domain.User, error) {
	var user domain.User
	err := u.Conn.Where("id = ?", id).First(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

// GetByEmail will get user data by given email
func (u *UserRepository) GetByEmail(c *gin.Context, email string) (domain.User, error) {
	var user domain.User
	err := u.Conn.Where("email = ?", email).First(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

// Store will store user data to database
func (u *UserRepository) Store(c *gin.Context, user *domain.User) error {
	err := u.Conn.Create(user).Error
	if err != nil {
		return err
	}
	return nil
}
