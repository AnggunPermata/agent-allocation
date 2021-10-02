package main

import (
	"fmt"

	"github.com/anggunpermata/agent-allocation/auth"
	"github.com/anggunpermata/agent-allocation/config"
	"github.com/anggunpermata/agent-allocation/routes"
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	config.InitDB()
	config.InitPort()
	auth.LogMiddlewares((e))
	routes.New(e)
	Port := fmt.Sprintf(":%d", config.PORT)
	if err := e.Start(Port); err != nil {
		e.Logger.Fatal(err)
	}
}
