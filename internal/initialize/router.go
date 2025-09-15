package initialize

import (
	"fmt"

	c "github.com/albertbui010/go-ecommerce-backend-api/internal/controller"
	"github.com/albertbui010/go-ecommerce-backend-api/internal/middlewares"
	"github.com/gin-gonic/gin"
)

func AA() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Before ---> AA")
		c.Next()
		fmt.Println("After ---> AA")
	}
}
func BB() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Before ---> BB")
		c.Next()
		fmt.Println("After ---> BB")
	}
}
func CC(c *gin.Context) {
	fmt.Println("Before ---> AA")
	c.Next()
	fmt.Println("After ---> AA")
}

func InitRouter() *gin.Engine {
	r := gin.Default()

	r.Use(middlewares.AuthenMiddleware(), AA(), BB(), CC)

	v1 := r.Group("/v1/2025") // Version 1
	{
		v1.GET("/ping", c.NewPongController().Pong)
		v1.GET("/user", c.NewUserController().GetUserById)
		// v1.PATCH("/ping", Pong)
		// v1.DELETE("/ping", Pong)
		// v1.OPTIONS("/ping", Pong)
		// v1.HEAD("/ping", Pong)
	}

	v2 := r.Group("/v2/2025") // Version 2
	{
		v2.GET("/ping", c.NewPongController().Pong)
		v2.GET("/user", c.NewUserController().GetUserById)
		// v2.PATCH("/ping", Pong)
		// v2.DELETE("/ping", Pong)
		// v2.OPTIONS("/ping", Pong)
		// v2.HEAD("/ping", Pong)
	}

	return r
}
