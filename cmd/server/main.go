package main

import (
	"github.com/albertbui010/go-ecommerce-backend-api/internal/routers"
)

func main() {
	r := routers.NewRouter()

	r.Run(":8081") // Default run on port 8080
}
