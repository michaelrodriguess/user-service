package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/michaelrodriguess/user-service/internal/model"
	"github.com/michaelrodriguess/user-service/internal/service"
)

type UserHandler struct {
	service *service.UserService
}

func NewUserHandler(service *service.UserService) *UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) CreateUserHandler(c *gin.Context) {
	var req model.UserRequest

	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	resp, err := h.service.CreateUserService(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		return
	}

	c.JSON(http.StatusCreated, resp)
}

func (h *UserHandler) GetAllAdminsUser(c *gin.Context) {

	resp, err := h.service.GetAllAdminsUser()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		return
	}

	c.JSON(http.StatusOK, resp)
}

func (h *UserHandler) GetAllUsers(c *gin.Context) {

	resp, err := h.service.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		return
	}

	c.JSON(http.StatusOK, resp)
}

func (h *UserHandler) DeleteUserHandler(c *gin.Context) {
	uuidUser := c.Query("uuid_user")
	if uuidUser == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "uuid_user is require."})

		return
	}

	err := h.service.DeleteUserByUUID(uuidUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting user. " + err.Error()})

		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User with uuid " + uuidUser + " deleted(SOFT) successfully"})

}
