package usecase

import (
	"Foundries/domain"
)

type UserUsecase struct {
	userRepo domain.UserRepository
}

// NewUserUsecase will create new an userUsecase object representation of domain.UserUsecase interface
func NewUserUsecase(userRepo domain.UserRepository) domain.UserUsecase {
	return &UserUsecase{userRepo}
}

// Fetch will fetch all user data from database
func (u *UserUsecase) Fetch() ([]domain.User, error) {
	users, err := u.userRepo.Fetch()
	if err != nil {
		return nil, err
	}
	return users, nil
}

// GetByID will get user data by given id
func (u *UserUsecase) GetByID(id string) (domain.User, error) {
	user, err := u.userRepo.GetByID(id)
	if err != nil {
		return user, err
	}
	return user, nil
}

// GetByEmail will get user data by given email
func (u *UserUsecase) GetByEmail(email string) (domain.User, error) {
	user, err := u.userRepo.GetByEmail(email)
	if err != nil {
		return user, err
	}
	return user, nil
}

// Store will store user data to database
func (u *UserUsecase) Store(user *domain.User) error {
	err := u.userRepo.Store(user)
	if err != nil {
		return err
	}
	return nil
}
