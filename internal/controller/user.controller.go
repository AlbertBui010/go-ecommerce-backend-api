package controller

import (
	"fmt"

	"github.com/albertbui010/go-ecommerce-backend-api/internal/service"
	"github.com/albertbui010/go-ecommerce-backend-api/internal/vo"
	"github.com/albertbui010/go-ecommerce-backend-api/pkg/response"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService service.IUserService
}

func NewUserController(
	userService service.IUserService,
) *UserController {
	return &UserController{
		userService: userService,
	}
}

func (uc *UserController) Register(c *gin.Context) {
	var params vo.UserRegistrationRequest
	if err := c.ShouldBindJSON(&params); err != nil {
		response.ErrorResponse(c, response.ErrCodeParamInvalid, "")
		return
	}

	fmt.Printf("Email params: %s", params.Email)

	result := uc.userService.Register(params.Email, params.Purpose)
	response.SuccessResponse(c, result, nil)
}
