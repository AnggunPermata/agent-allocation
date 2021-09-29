package main

import (
	"fmt"

	"github.com/anggunpermata/agent-allocation/config"
	"github.com/anggunpermata/agent-allocation/routes"
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	config.InitDB()
	config.InitPort()
	//auth.LogMiddlewares((e))
	routes.New(e)
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", config.PORT)))
}
