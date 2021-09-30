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

	//check if customer has already initiate a channel and still active
	statusChat, err := database.GetOneChannelById(customerId)
	//if channel id > 0 then the channel already exist
	if statusChat.ID > 0 {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message":              "You have initiated a channel",
			"Initiated Channel ID": statusChat.ID,
		})
	}

	availableAgent, err := database.GetOneAgent(c)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "cannot get agent account",
		})
	}

	//update count_active_channel into agents table
	// availableAgent.ID = availableAgent.ID
	// availableAgent.Username = availableAgent.Username
	// availableAgent.Agent_Status = availableAgent.Agent_Status
	// availableAgent.Token = availableAgent.Token
	// availableAgent.Password = availableAgent.Password
	availableAgent.Count_Active_Channel = availableAgent.Count_Active_Channel + 1
	c.Bind(&availableAgent)
	updateAgent, err2 := database.UpdateAgent(availableAgent)
	if err2 != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "cannot update agent data",
		})
	}

	newChannel.CustomerID = uint(customerId)
	newChannel.AgentID = updateAgent.ID
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
