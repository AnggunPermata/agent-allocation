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
	// Agent login & logout
	e.POST("/agent/login", controller.AgentLogin)
	eJwt.PUT("/agent/:agent_id/logout", controller.AgentLogout)

	//agent sends message
	eJwt.POST("/agent/:agent_id/chat/send", controller.AgentAsSender)

	//agent see messages
	eJwt.GET("agent/:agent_id/chat", controller.AgentGetAllChannelMessages)

	//agent resolved message
	eJwt.POST("agent/:agent_id/chat/resolve", controller.AgentResolveChat)

	// Customer login & logout
	e.POST("customer/login", controller.CustomerLogin)
	eJwt.PUT("customer/:customer_id/logout", controller.CustomerLogout)

	//customer initiate new channel to chat with agent
	eJwt.POST("customer/:customer_id/chat/initiate", controller.NewChannel)

	//customer sends message
	eJwt.POST("customer/:customer_id/chat/send", controller.CustomerAsSender)

	//customer see messages
	eJwt.GET("customer/:customer_id/chat", controller.CustomerGetAllChannelMessages)
}
