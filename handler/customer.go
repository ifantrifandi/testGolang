package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"trainingapi/auth"
	"trainingapi/customer"
	"trainingapi/helper"
)

type customerHandler struct {
	customerService customer.Service
	authService auth.Service
}

func NewCustomerHandler(customerService customer.Service, authService auth.Service)  *customerHandler {
	return &customerHandler{customerService, authService}
}

func (h *customerHandler) RegisterCustomer(c *gin.Context) {
	var input customer.RegisterCustomerInput

	err := c.ShouldBindJSON(&input)

	if err != nil {
		errors := helper.FormatError(err)
		errorMessage := gin.H{"errors" : errors}
		response := helper.APIResponse("Invalid Data", http.StatusUnprocessableEntity, "failed", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newCustomer , err := h.customerService.RegisterCustomer(input)

	if err != nil {
		response := helper.APIResponse("Register Account Failed", http.StatusBadRequest, "failed", err)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := customer.FormatCustomer(newCustomer, "")

	response := helper.APIResponse("Account has been registered", 201, "success", formatter)

	c.JSON(http.StatusCreated, response)
}

func (h *customerHandler) LoginCustomer(c *gin.Context) {
	var input customer.LoginCustomerInput

	err := c.ShouldBindJSON(&input)

	if err != nil {
		errors := helper.FormatError(err)
		errorMessage := gin.H{"errors" : errors}
		response := helper.APIResponse("Invalid Email / Password", http.StatusBadRequest, "failed", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	customerLogin, err := h.customerService.Login(input)
	fmt.Println("handler")
	if err != nil {
		var message map[string]string
		message = map[string]string{}
		message["error"] = err.Error()
		response := helper.APIResponse("Invalid Email / Password", http.StatusBadRequest, "failed", message)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	token, err := h.authService.GenerateToken(customerLogin.ID)

	if err != nil {
		errors := helper.FormatError(err)
		errorMessage := gin.H{"errors" : errors}
		response := helper.APIResponse("Invalid Email / Password", http.StatusBadRequest, "failed", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := customer.FormatCustomer(customerLogin, token)

	response := helper.APIResponse("Login Successful", 200, "success", formatter)

	c.JSON(http.StatusOK, response)
}

func (h *customerHandler) FindCustomer(c *gin.Context) {

	customers := h.customerService.FindCustomer()

	formatter := customer.DataFormatCustomer(customers)

	c.JSON(http.StatusOK, formatter)
}

func (h *customerHandler) FindCustomerByName(c *gin.Context) {
	name := c.Param("name")
	customers := h.customerService.FindCustomerByName(name)
	c.JSON(http.StatusOK , customers)
}

func (h *customerHandler) FindCustomerById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response := helper.APIResponse("Please enter valid ID", http.StatusUnprocessableEntity, "error", nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	findCustomer := h.customerService.FindCustomerById(int(id))

	formatter := customer.FormatCustomer(findCustomer, "haiya")

	response := helper.APIResponse("ID Found", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)

}