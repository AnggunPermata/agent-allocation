package controller

import (
	"net/http"

	"github.com/anggunpermata/agent-allocation/auth"
	"github.com/anggunpermata/agent-allocation/database"
	"github.com/anggunpermata/agent-allocation/models"
	"github.com/labstack/echo"
)

//To check user's authorization by using user Id

func AuthorizedAgent(c echo.Context) bool {
	_, role := auth.ExtractTokenUserId(c)
	if role != "agent" {
		return false
	}
	return true
}

func AgentLogin(c echo.Context) error {
	inputData := models.Login_Form{}
	c.Bind(&inputData)
	userData := models.Agent{
		Username: inputData.Username,
		Password: inputData.Password,
	}
	c.Bind(&userData)

	agent, err := database.AgentLogin(userData.Username, userData.Password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "please check your username and password again.",
		})
	}

	showAgentData := map[string]interface{}{
		"ID":       agent.ID,
		"Username": "@" + agent.Username,
		"Token":    agent.Token,
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":   "Hello, Let's Start!",
		"user-data": showAgentData,
	})
}
