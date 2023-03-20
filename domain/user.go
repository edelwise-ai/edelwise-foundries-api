package domain

import (
	"time"
)

type User struct {
	ID        string    `json:"id"`
	Email     string    `json:"email" gorm:"unique"`
	Password  string    `json:"password"`
	Verified  bool      `json:"verified" gorm:"default:false"`
	CreatedAt time.Time `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	//Datasets  []Dataset `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	//Projects  []Project `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type UserUsecase interface {
	Fetch() ([]User, error)
	GetByID(id string) (User, error)
	GetByEmail(email string) (User, error)
	Store(user *User) error
}

type UserRepository interface {
	Fetch() ([]User, error)
	GetByID(id string) (User, error)
	GetByEmail(email string) (User, error)
	Store(user *User) error
}
