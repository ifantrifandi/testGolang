package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"trainingapi/user"
)

type userHandler struct {
	userService user.Service
}

func NewUserHandler(userService user.Service) *userHandler {
	return &userHandler{userService}
}

func (h *userHandler) FindUser(c *gin.Context) {
	var response []user.User

	response = h.userService.FindUser()

	c.JSON(http.StatusOK, response)
}