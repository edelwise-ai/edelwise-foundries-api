package repository

import "Foundries/users/models"

type UserRepository interface {
	Fetch() ([]models.User, error)
	GetByID(id string) (models.User, error)
	GetByEmail(email string) (models.User, error)
	Store(user *models.User) error
}
