package routes

import (
	"github.com/anggunpermata/agent-allocation/constant"
	"github.com/anggunpermata/agent-allocation/controller"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func New(e *echo.Echo) {
	eJwt := e.Group("")
	eJwt.Use(middleware.JWT([]byte(constant.SECRET_JWT)))
	// Agent login
	e.POST("agent/login", controller.AgentLogin)
	// Customer login
	e.POST("customer/login", controller.CustomerLogin)

}
