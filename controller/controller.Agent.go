package controller

import (
	"net/http"
	"strconv"

	"github.com/anggunpermata/agent-allocation/auth"
	"github.com/anggunpermata/agent-allocation/database"
	"github.com/anggunpermata/agent-allocation/models"
	"github.com/labstack/echo"
)

//To check user's authorization by using user Id

func AuthorizedAgent(agentId int, c echo.Context) error {
	// _, role := auth.ExtractTokenUserId(c)
	// if role != "agent" {
	// 	return false
	// }
	// return true

	agent, err := database.GetOneAgentById(agentId)
	loggedInAgentId, role := auth.ExtractTokenUserId(c)
	if loggedInAgentId != int(agent.ID) || err != nil || role != "agent" {
		return echo.NewHTTPError(http.StatusUnauthorized, "Cannot access this account")
	}
	return nil
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

func AgentResolveChat(c echo.Context) error {
	//update the messages table.chat_status into "resolved"
	//update the channels table.chat_status into "resolved"
	// minus -1 for count_active_channel in agent table

	agentId, err := strconv.Atoi(c.Param("agent_id"))
	customerId, err := strconv.Atoi(c.FormValue("customer_id"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "please check id again.",
		})
	}

	if err := AuthorizedAgent(int(agentId), c); err != nil {
		return err
	}

	minusCount, err := database.GetOneAgentById(agentId)
	minusCount.Count_Active_Channel = minusCount.Count_Active_Channel - 1
	c.Bind(&minusCount)
	_, err = database.UpdateAgent(minusCount)

	resolveChannel, err2 := database.GetOneChannelByAgentId(agentId, customerId)
	if err2 != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "cannot find active channel",
		})
	}
	resolveChannel.Chat_Status = "resolved"
	c.Bind(&resolveChannel)
	resolved, err3 := database.UpdateChannel(resolveChannel)
	if err3 != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "cannot change chat_status",
		})
	}

	resolveLastMessage, err := database.GetLastMessageByChannelId(int(resolved.ID))
	resolveLastMessage.Chat_Status = "resolved"
	c.Bind(&resolveLastMessage)
	_, err = database.UpdateChatStatus(resolveLastMessage)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "cannot update chat status",
		})
	}

	mapData := map[string]interface{}{
		"Channel ID":  resolved.ID,
		"Customer ID": resolved.CustomerID,
		"Agent ID":    resolved.AgentID,
		"Chat Status": resolved.Chat_Status,
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"Message": "Success",
		"Data":    mapData,
	})
}

func AgentLogout(c echo.Context) error {
	agentId, err := strconv.Atoi(c.Param("agent_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid id",
		})
	}
	if err = AuthorizedAgent(agentId, c); err != nil {
		return err
	}
	logout, err := database.GetOneAgentById(agentId)
	logout.Token = ""
	logout.Agent_Status = "unactive"
	c.Bind(&logout)
	agent, err := database.UpdateAgent(logout)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "cannot logout",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":  "Make sure to online again at 9 am!",
		"Agent ID": agent.ID,
		"Username": agent.Username,
	})
}
