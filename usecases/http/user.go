package http

import (
	"Foundries/models"
	"Foundries/repositories"
	"Foundries/usecases"
)

type UserUsecase struct {
	userRepo repositories.UserRepository
}

// NewUserUsecase will create new an userUsecase object representation of models.UserUsecase interface
func NewUserUsecase(userRepo repositories.UserRepository) usecases.UserUsecase {
	return &UserUsecase{userRepo}
}

// Fetch will fetch all user data from database
func (u *UserUsecase) Fetch() ([]models.User, error) {
	users, err := u.userRepo.Fetch()
	if err != nil {
		return nil, err
	}
	return users, nil
}

// GetByID will get user data by given id
func (u *UserUsecase) GetByID(id string) (models.User, error) {
	user, err := u.userRepo.GetByID(id)
	if err != nil {
		return user, err
	}
	return user, nil
}

// GetByEmail will get user data by given email
func (u *UserUsecase) GetByEmail(email string) (models.User, error) {
	user, err := u.userRepo.GetByEmail(email)
	if err != nil {
		return user, err
	}
	return user, nil
}

// Store will store user data to database
func (u *UserUsecase) Store(user *models.User) error {
	err := u.userRepo.Store(user)
	if err != nil {
		return err
	}
	return nil
}
