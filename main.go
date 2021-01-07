package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"trainingapi/customer"
	"trainingapi/handler"
	"trainingapi/user"
)

func main() {

	dsn := "host=localhost user=toktok2 password='' dbname=todo port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	customerRepository := customer.NewRepository(db)
	customerService := customer.NewService(customerRepository)
	customerHandler := handler.NewCustomerHandler(customerService)

	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)
	userHandler := handler.NewUserHandler(userService)

	router := gin.Default()

	api := router.Group("/api/v1")

	api.GET("/customer", customerHandler.FindCustomer)
	api.POST("/customer", customerHandler.RegisterCustomer)
	api.GET("/user", userHandler.FindUser)
	//":portNumber" -> inside Run() if wanna define Port, default 8080
	newErr := router.Run()
	if newErr != nil {
		fmt.Println("error nih")
	}
}
