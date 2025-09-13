package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/v1/2025") // Version 1
	{
		v1.GET("/ping/:name", Pong)
		v1.PUT("/ping", Pong)
		v1.PATCH("/ping", Pong)
		v1.DELETE("/ping", Pong)
		v1.OPTIONS("/ping", Pong)
		v1.HEAD("/ping", Pong)
	}

	v2 := r.Group("/v2/2025") // Version 2
	{
		v2.GET("/ping/:name", Pong)
		v2.PUT("/ping", Pong)
		v2.PATCH("/ping", Pong)
		v2.DELETE("/ping", Pong)
		v2.OPTIONS("/ping", Pong)
		v2.HEAD("/ping", Pong)
	}

	return r
}

func Pong(ctx *gin.Context) {
	name := ctx.Param("name")
	uid := ctx.Query("uid")

	ctx.JSON(http.StatusOK, gin.H{
		"message": "pong...ping" + name,
		"name":    name,
		"uid":     uid,
	})
}
