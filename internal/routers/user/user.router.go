package user

import (
	"github.com/albertbui010/go-ecommerce-backend-api/internal/wire"
	"github.com/gin-gonic/gin"
)

type UserRouter struct{}

func (pr *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	// DI
	userController, _ := wire.InitUserRouterHandler()

	// public router
	userRouterPublic := Router.Group("/user")
	{
		userRouterPublic.POST("/register", userController.Register)
		userRouterPublic.POST("/otp")
	}
	// private router

	// userRouterPrivate := Router.Group("/user")
	// // userRouterPrivate.Use(Limiter())
	// // userRouterPrivate.Use(Auth())
	// // userRouterPrivate.Use(Permission())
	// {
	// 	userRouterPrivate.POST("/register")
	// 	userRouterPrivate.POST("/otp")
	// }
}
