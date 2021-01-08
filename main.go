package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"trainingapi/auth"
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
	authService := auth.NewService()
	customerHandler := handler.NewCustomerHandler(customerService, authService)


	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)
	userHandler := handler.NewUserHandler(userService)

	router := gin.Default()

	api := router.Group("/api/v1")

	api.POST("/customer", customerHandler.RegisterCustomer)
	api.POST("/customer/login", customerHandler.LoginCustomer)
	api.GET("/customer", customerHandler.FindCustomer)
	api.GET("/customer/id/:id", customerHandler.FindCustomerById)
	api.GET("/customer/name/:name", customerHandler.FindCustomerByName)
	api.GET("/user", userHandler.FindUser)

	//":portNumber" -> inside Run() if wanna define Port, default 8080
	newErr := router.Run()
	if newErr != nil {
		fmt.Println("error nih")
	}
}

// for Authorization middleware

//func authMiddleware(c *gin.Context){
//	authHeader := c.GetHeader("Authorization")
//
//	if !strings.Contains(authHeader, "Bearer") {
//		response := helper.APIResponse("Unathorized", http.StatusUnauthorized, "error", nil)
//		c.AbortWithStatusJSON(http.StatusUnauthorized, response)
//
//		return
//	}
//
//	authToken := strings.Split(authHeader, " ")
//	tokenString := ""
//
//	if len(authToken) == 2 {
//		tokenString = authToken[1]
//	}
//
//	token, err :=
//
//}
