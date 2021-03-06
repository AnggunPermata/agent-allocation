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

	if err := AuthorizedCustomer(int(customerId), c); err != nil {
		return err
	}

	status, err := database.GetOneCustomerById(customerId)
	if status.Token == "" {
		return c.JSON(http.StatusBadRequest, "You have to login again")
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

	if err := AuthorizedAgent(int(agentId), c); err != nil {
		return err
	}

	status, err := database.GetOneAgentById(agentId)
	if status.Token == "" {
		return c.JSON(http.StatusBadRequest, "You have to login again")
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

//this function used by agent to see the chats from the first time it was
// initiated by the customer
func AgentGetAllChannelMessages(c echo.Context) error {

	data := models.Check_All_Message_Input{}
	c.Bind(&data)

	agentId := data.AgentID
	customerId := data.CustomerID
	channelId := data.ChannelID

	if err := AuthorizedAgent(int(agentId), c); err != nil {
		return err
	}

	status, _ := database.GetOneAgentById(int(agentId))
	if status.Token == "" {
		return c.JSON(http.StatusBadRequest, "You have to login again")
	}

	check, _ := database.GetOneChannelByAgentAndCustomerId(int(agentId), int(customerId))
	if check.ID != uint(channelId) {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "you are not allowed to enter this channel/ channel is not exist",
		})
	}

	msg, err4 := database.GetChannelDataById(int(channelId))
	if err4 != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "cannot find message data",
		})
	}

	var allMessage []map[string]interface{}
	for i := 0; i < len(msg); i++ {
		mapMsg := map[string]interface{}{
			"Message ID":                    msg[i].ID,
			"Channel ID":                    msg[i].ChannelID,
			"Sender Role / Sender ID":       msg[i].Sender_Role + " / " + strconv.Itoa(int(msg[i].SenderID)),
			"Recipient Role / Recipient ID": msg[i].Recipient_Role + " / " + strconv.Itoa(int(msg[i].RecipientID)),
			"Text Message":                  msg[i].TextMessage,
			"Chat_Status":                   msg[i].Chat_Status,
			"Arrived at":                    msg[i].CreatedAt,
		}
		allMessage = append(allMessage, mapMsg)
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data":    allMessage,
	})
}

//this function used by agent to see the chats from the first time it was
// initiated by the customer
func CustomerGetAllChannelMessages(c echo.Context) error {
	customerId, err := strconv.Atoi(c.Param("customer_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid id",
		})
	}

	if err := AuthorizedCustomer(int(customerId), c); err != nil {
		return err
	}

	status, err := database.GetOneCustomerById(customerId)
	if status.Token == "" {
		return c.JSON(http.StatusBadRequest, "You have to login again")
	}

	channelId, err2 := strconv.Atoi(c.FormValue("channel_id"))
	if err2 != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid channel id",
		})
	}

	check, err := database.GetOneChannelById(customerId)
	if check.ID != uint(channelId) {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "you are not allowed to enter this channel/ channel is not exist",
		})
	}

	msg, err4 := database.GetChannelDataById(channelId)
	if err4 != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "cannot find message data",
		})
	}

	var allMessage []map[string]interface{}
	for i := 0; i < len(msg); i++ {
		mapMsg := map[string]interface{}{
			"Message ID":                    msg[i].ID,
			"Channel ID":                    msg[i].ChannelID,
			"Sender Role / Sender ID":       msg[i].Sender_Role + " / " + strconv.Itoa(int(msg[i].SenderID)),
			"Recipient Role / Recipient ID": msg[i].Recipient_Role + " / " + strconv.Itoa(int(msg[i].RecipientID)),
			"Text Message":                  msg[i].TextMessage,
			"Chat_Status":                   msg[i].Chat_Status,
			"Arrived at":                    msg[i].CreatedAt,
		}
		allMessage = append(allMessage, mapMsg)
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data":    allMessage,
	})
}
