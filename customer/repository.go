package customer

import (
	"gorm.io/gorm"
)

type Repository interface{
	Save(customer Customer) (Customer, error)
	FindCustomer() []Customer
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(customer Customer) (Customer, error){
	err := r.db.Create(&customer).Error

	if err != nil {
		return customer, err
	}

	return customer, nil
}

func (r *repository) FindCustomer() []Customer {

	var customers []Customer

	r.db.Find(&customers)

	return customers
}