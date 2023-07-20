package usercontrollers

import "gorm.io/gorm"

type UserRepository interface {
	FindAll() ([]User, error)
	FindByUnique(Unique string) (User, error)
	Create(user User) (User, error)
}

type Repository struct {
	db *gorm.DB
}

func NewRepositoryUser(db *gorm.DB) *Repository {
	return &Repository{db}
}
