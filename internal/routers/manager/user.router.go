package manager

import (
	"github.com/albertbui010/go-ecommerce-backend-api/internal/wire"
	"github.com/gin-gonic/gin"
)

type UserRouter struct{}

func (pr *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	// Non-dependency
	// ur := repo.NewUserRepository()
	// us := service.NewUserService(ur)
	// userHandlerNonDependency := controller.NewUserController(us)
	userController, _ := wire.InitUserRouterHandler()
	// WIRE go

	// public router
	userRouterPublic := Router.Group("/admin/user")
	{
		userRouterPublic.POST("/register", userController.Register)
		// userRouterPublic.POST("/otp")
	}
	// private router

	userRouterPrivate := Router.Group("/admin/user")
	// userRouterPrivate.Use(Limiter())
	// userRouterPrivate.Use(Auth())
	// userRouterPrivate.Use(Permission())
	{
		userRouterPrivate.POST("/active_user")
	}
}
