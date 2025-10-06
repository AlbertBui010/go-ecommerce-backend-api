//go:build wireinject

package wire

import (
	"github.com/albertbui010/go-ecommerce-backend-api/internal/controller"
	"github.com/albertbui010/go-ecommerce-backend-api/internal/repo"
	"github.com/albertbui010/go-ecommerce-backend-api/internal/service"
	"github.com/google/wire"
)

func InitUserRouterHandler() (*controller.UserController, error) {
	wire.Build(
		// Declare dependencies here
		repo.NewUserRepository,
		repo.NewUserAuthRepository,
		service.NewUserService,
		controller.NewUserController,
	)

	return new(controller.UserController), nil
}
