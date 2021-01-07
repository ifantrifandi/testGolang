package user

import (
	"gorm.io/gorm"
)

type Repository interface {
	//Save(user User) (User error)
	Find() []User
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Find() []User {

	var users []User

	r.db.Find(&users)

	return users
}