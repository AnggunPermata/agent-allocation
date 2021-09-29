package controller

import (
	"net/http"
	"strconv"

	"github.com/anggunpermata/agent-allocation/constant"
	"github.com/anggunpermata/agent-allocation/database"
	"github.com/anggunpermata/agent-allocation/models"
	"github.com/labstack/echo"
)

//Customer requesting a channel to chat with the agent
func NewChannel(c echo.Context) error {
	newChannel := models.Channel{}
	customerId, err := strconv.Atoi(c.Param("customer_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Sorry, curtomer id not exist/ you dont have the authorization to access the account")
	}
	availableAgent, err := database.GetOneAgent(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "cannot get agent account",
		})
	}

	newChannel.CustomerID = uint(customerId)
	newChannel.AgentID = availableAgent.ID
	newChannel.Chat_Status = "active"

	c.Bind(&newChannel)
	created, err := database.CreateNewChannel(newChannel)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "cannot create new channel",
		})
	}

	channelData := map[string]interface{}{
		"Channel ID":  created.ID,
		"Customer ID": created.CustomerID,
		"Agent ID":    created.AgentID,
		"Chat Status": "in queue",
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": constant.Default_Welcome,
		"data":    channelData,
	})
}
