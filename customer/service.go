package customer

import "golang.org/x/crypto/bcrypt"

type Service interface {
	RegisterCustomer(input RegisterCustomerInput) (Customer, error)
	FindCustomer() []Customer
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func(s *service) RegisterCustomer (input RegisterCustomerInput) (Customer, error) {
	customer := Customer{}
	customer.Name = input.Name
	newPasswordHash , err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil {
		return customer, err
	}
	customer.PasswordHash = string(newPasswordHash)
	customer.Role = "user"
	customer.Occupation = input.Occupation

	newCustomer, err := s.repository.Save(customer)

	if err != nil {
		return newCustomer , err
	}

	return newCustomer, nil
}

func (s *service) FindCustomer() []Customer {

	customers := s.repository.FindCustomer()

	return  customers
}