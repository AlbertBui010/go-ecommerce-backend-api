package manager

import "github.com/gin-gonic/gin"

type AdminRouter struct{}

func (pr *AdminRouter) InitAdminRouter(Router *gin.RouterGroup) {
	// public router
	adminRouterPublic := Router.Group("/admin")
	{
		adminRouterPublic.POST("/login")
	}
	// private router

	// adminRouterPrivate := Router.Group("/admin/user")
	// // adminRouterPrivate.Use(Limiter())
	// // adminRouterPrivate.Use(Auth())
	// // adminRouterPrivate.Use(Permission())
	// {
	// 	adminRouterPrivate.POST("/active_user")
	// }
}
