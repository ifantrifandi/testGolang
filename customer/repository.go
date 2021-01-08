package customer

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
)

type Repository interface{
	Save(customer Customer) (Customer, error)
	FindCustomer() []Customer
	FindCustomerByName(name string) []Customer
	FindCustomerById(id int) Customer
	CheckEmailAndPassword(email string, password string) (Customer, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

//setiap membuat method untuk cari di db, biasakan menggunakan err := gorm query (CRUD).Error => untuk mengetahui apakah terdapat error apa tidak saat menggunakan gorm query.
//sehingga pastikan nantinya setiap repository akan mengembalikan 2 data => 1 data hasil db 1 data apabila terjadi error
//disini bnyk yang ga pake buat pembelajaran aja, best practicenya menggunakan var err untuk mengecek apakah ada error atau tidak sehingga data yang dikembalikan nantinya tidak aneh

func (r *repository) Save(customer Customer) (Customer, error){
	err := r.db.Create(&customer).Error

	if err != nil {
		return customer, err
	}

	return customer, nil
}

func (r *repository) FindCustomer() []Customer {

	var customers []Customer

	//err := r.db.Find(&customers).Error
	r.db.Find(&customers)

	return customers
}

func (r *repository) FindCustomerByName(name string) []Customer{
	var customers []Customer
	var findCustomer = "%" + name + "%"
	//err := r.db.Where("name like ?", findCustomer).Find(&customers)
	r.db.Where("name Like ?", findCustomer).Find(&customers)

	return customers
}

func (r *repository) FindCustomerById(id int) Customer {
	var customer Customer

	r.db.Where("id = ?", id).Find(&customer)

	return customer
}

func (r *repository) CheckEmailAndPassword(name string, password string) (Customer, error) {

	var customer Customer
	err := r.db.Where("name = ?", name).First(&customer).Error
	fmt.Println("Kepanggil")
	if err != nil {
		fmt.Println("Kepanggil error")
		return customer, errors.New("Invalid Email / Password")
	}

	//if err != nil {
	//	return customer, errors.New("hihihaha")
	//	//a := errors.Is(err, gorm.ErrRecordNotFound)
	//	//if a {
	//	//	return customer, errors.New("Invalid Email / Password")
	//	//}
	//}
	//return customer, nil
	//comparePassword := bcrypt.CompareHashAndPassword([]byte(customer.PasswordHash), []byte(password))
	//fmt.Println(comparePassword)
	//
	//if comparePassword != nil {
	//	return customer, errors.New("Invalid Email / Password")
	//}
	//
	//return customer, nil
	return customer, nil
}