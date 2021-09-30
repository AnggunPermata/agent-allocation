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

	//agent sends message
	e.POST("agent/:agent_id/chat/send", controller.AgentAsSender)

	// Customer login
	e.POST("customer/login", controller.CustomerLogin)

	//customer initiate new channel to chat with agent
	e.POST("customer/:customer_id/chat/initiate", controller.NewChannel)

	//customer sends message
	e.POST("customer/:customer_id/chat/send", controller.CustomerAsSender)
}
