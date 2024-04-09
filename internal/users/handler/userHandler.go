package handler

import (
	"Project1/internal/users/model"
	"Project1/internal/users/service"
	handler "Project1/pkg/response"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service *service.UserService
}

func NewUserHandler(service *service.UserService) *UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) SignUp(c *gin.Context) {
	var input model.User

	if err := c.BindJSON(&input); err != nil {
		handler.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.service.CreateUser(input)
	if err != nil {
		handler.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *UserHandler) SignIn(c *gin.Context) {
	var input model.User

	if err := c.BindJSON(&input); err != nil {
		handler.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	token, err := h.service.GenerateToken(input.Username, input.Password)
	if err != nil {
		handler.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}

func (h *UserHandler) DeleteUser(c *gin.Context) {
	userId := c.Param("id")
	id, err := strconv.Atoi(userId)
	if err != nil {
		handler.NewErrorResponse(c, http.StatusBadRequest, "Invalid user ID")
		return
	}

	if err := h.service.DeleteUser(id); err != nil {
		handler.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
