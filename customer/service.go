package customer

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	RegisterCustomer(input RegisterCustomerInput) (Customer, error)
	FindCustomer() []Customer
	FindCustomerByName(name string) []Customer
	FindCustomerById(id int) Customer
	Login(input LoginCustomerInput) (Customer, error)
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

func (s *service) FindCustomerByName(name string) []Customer {
	customers := s.repository.FindCustomerByName(name)

	return customers
}

func (s *service) FindCustomerById(id int) Customer{
	customer := s.repository.FindCustomerById(id)

	return customer
}

func (s *service) Login(input LoginCustomerInput) (Customer, error) {
	customer , err := s.repository.CheckEmailAndPassword(input.Name, input.Password)
	fmt.Println("service")
	if err != nil {
		fmt.Println("error service")
		return customer, err
	}

	return customer, nil
}