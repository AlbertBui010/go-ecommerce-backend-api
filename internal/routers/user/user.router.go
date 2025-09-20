package user

import "github.com/gin-gonic/gin"

type UserRouter struct{}

func (pr *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	// public router
	userRouterPublic := Router.Group("/user")
	{
		userRouterPublic.POST("/register")
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
