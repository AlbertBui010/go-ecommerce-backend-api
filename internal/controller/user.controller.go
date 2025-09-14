package controller

import (
	"net/http"

	"github.com/albertbui010/go-ecommerce-backend-api/internal/service"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService *service.UserService
}

func NewUserController() *UserController {
	return &UserController{
		userService: service.NewUserService(),
	}
}

func (uc *UserController) GetUserById(ctx *gin.Context) {
	user := uc.userService.GetUserById()

	ctx.JSON(http.StatusOK, gin.H{
		"user": "user object " + user,
	})
}
