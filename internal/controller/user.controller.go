package controller

import (
	"github.com/albertbui010/go-ecommerce-backend-api/internal/service"
	"github.com/albertbui010/go-ecommerce-backend-api/pkg/response"
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

	response.SuccessResponse(ctx, 20001, user)
}
