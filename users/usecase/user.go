package usecase

import (
	"Foundries/domain"
	"github.com/gin-gonic/gin"
)

type UserUsecase struct {
	userRepo domain.UserRepository
}

// NewUserUsecase will create new an userUsecase object representation of domain.UserUsecase interface
func NewUserUsecase(userRepo domain.UserRepository) domain.UserUsecase {
	return &UserUsecase{userRepo}
}

// Fetch will fetch all user data from database
func (u *UserUsecase) Fetch(c *gin.Context) ([]domain.User, error) {
	users, err := u.userRepo.Fetch(c)
	if err != nil {
		return nil, err
	}
	return users, nil
}

// GetByID will get user data by given id
func (u *UserUsecase) GetByID(c *gin.Context, id string) (domain.User, error) {
	user, err := u.userRepo.GetByID(c, id)
	if err != nil {
		return user, err
	}
	return user, nil
}

// GetByEmail will get user data by given email
func (u *UserUsecase) GetByEmail(c *gin.Context, email string) (domain.User, error) {
	user, err := u.userRepo.GetByEmail(c, email)
	if err != nil {
		return user, err
	}
	return user, nil
}

// Store will store user data to database
func (u *UserUsecase) Store(c *gin.Context, user *domain.User) error {
	err := u.userRepo.Store(c, user)
	if err != nil {
		return err
	}
	return nil
}
