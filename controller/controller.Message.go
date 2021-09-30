package controller

import (
	"net/http"
	"strconv"

	"github.com/anggunpermata/agent-allocation/database"
	"github.com/anggunpermata/agent-allocation/models"
	"github.com/labstack/echo"
)

func CustomerAsSender(c echo.Context) error {
	//to get customer id using param
	customerId, err := strconv.Atoi(c.Param("customer_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid id",
		})
	}

	data := models.Input_Message{}
	c.Bind(&data)

	channelData, _ := database.GetOneChannelById(customerId)

	if channelData.ID < 1 {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "channel is not available.",
		})
	}

	//if channel has been resolved by the agent
	if channelData.Chat_Status == "resolved" {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Please initiate a new channel",
		})
	}

	saveData := models.Message{}
	saveData.ChannelID = channelData.ID
	saveData.Sender_Role = "customer"
	saveData.SenderID = uint(customerId)
	saveData.Recipient_Role = "agent"
	saveData.RecipientID = channelData.AgentID
	saveData.TextMessage = data.TextMessage
	saveData.Chat_Status = channelData.Chat_Status

	c.Bind(&saveData)

	//save the new message into database
	send, err := database.CreateNewMessage(saveData)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Cannot send message.",
		})
	}

	showMessageData := map[string]interface{}{
		"Message ID":          send.ID,
		"Channel ID":          send.ChannelID,
		"Sender Role / ID":    send.Sender_Role + " / " + strconv.Itoa(int(send.SenderID)),
		"Recipient Role / ID": send.Recipient_Role + " / " + strconv.Itoa(int(send.RecipientID)),
		"Text Message":        send.TextMessage,
		"Chat Status":         send.Chat_Status,
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success sending a new message.",
		"data":    showMessageData,
	})
}

func AgentAsSender(c echo.Context) error {
	//to get customer id using param
	agentId, err := strconv.Atoi(c.Param("agent_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid id",
		})
	}

	data := models.Input_Message{}
	c.Bind(&data)

	channelData, _ := database.GetOneChannelByAgentId(agentId, int(data.RecipientID))

	if channelData.ID < 1 {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "channel is not available.",
		})
	}

	//if channel has been resolved by the agent
	if channelData.Chat_Status == "resolved" {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Please wait until customer initiate new channel",
		})
	}

	saveData := models.Message{}
	saveData.ChannelID = channelData.ID
	saveData.Sender_Role = "agent"
	saveData.SenderID = uint(agentId)
	saveData.Recipient_Role = "customer"
	saveData.RecipientID = channelData.CustomerID
	saveData.TextMessage = data.TextMessage
	saveData.Chat_Status = channelData.Chat_Status

	c.Bind(&saveData)

	//save the new message into database
	send, err := database.CreateNewMessage(saveData)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Cannot send message.",
		})
	}

	showMessageData := map[string]interface{}{
		"Message ID":          send.ID,
		"Channel ID":          send.ChannelID,
		"Sender Role / ID":    send.Sender_Role + " / " + strconv.Itoa(int(send.SenderID)),
		"Recipient Role / ID": send.Recipient_Role + " / " + strconv.Itoa(int(send.RecipientID)),
		"Text Message":        send.TextMessage,
		"Chat Status":         send.Chat_Status,
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success sending a new message.",
		"data":    showMessageData,
	})
}
