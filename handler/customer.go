package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"trainingapi/customer"
	"trainingapi/helper"
)

type customerHandler struct {
	customerService customer.Service
}

func NewCustomerHandler(customerService customer.Service)  *customerHandler {
	return &customerHandler{customerService}
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

	formatter := customer.FormatCustomer(newCustomer, "tokennnnn")

	response := helper.APIResponse("Account has been registered", 201, "success", formatter)

	c.JSON(http.StatusCreated, response)
}

func (h *customerHandler) FindCustomer(c *gin.Context) {

	customers := h.customerService.FindCustomer()

	formatter := customer.DataFormatCustomer(customers)

	c.JSON(http.StatusOK, formatter)
}