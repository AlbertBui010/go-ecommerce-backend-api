package routers

import (
	"github.com/albertbui010/go-ecommerce-backend-api/internal/routers/manager"
	"github.com/albertbui010/go-ecommerce-backend-api/internal/routers/user"
)

type RouterGroup struct {
	User    user.UserGroupRouter
	Manager manager.ManagerGroupRouter
}

var RouterGroupApp = new(RouterGroup)
