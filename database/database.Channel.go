package database

import (
	"github.com/anggunpermata/agent-allocation/config"
	"github.com/anggunpermata/agent-allocation/models"
)

func CreateNewChannel(channel models.Channel) (models.Channel, error) {
	if err := config.DB.Save(&channel).Error; err != nil {
		return channel, err
	}
	return channel, nil
}

func GetChannelDataById(channelId int) ([]models.Message, error) {
	var channelData []models.Message
	if err := config.DB.Where("channel_id =?", channelId).Find(&channelData).Error; err != nil {
		return channelData, err
	}

	return channelData, nil
}

func GetOneChannelById(customer_Id int) (models.Channel, error) {
	var channel models.Channel
	if err := config.DB.Where("customer_id = ? AND chat_status =?", customer_Id, "active").Find(&channel).Error; err != nil {
		return channel, err
	}

	return channel, nil
}

func GetAllChannelByAgentId(agent_Id int) ([]models.Channel, error) {
	var channel []models.Channel
	if err := config.DB.Where("agent_id = ? AND chat_status =?", agent_Id, "active").Find(&channel).Error; err != nil {
		return channel, err
	}
	return channel, nil
}

func GetOneChannelByAgentId(agent_Id, customer_Id int) (models.Channel, error) {
	var channel models.Channel
	if err := config.DB.Where("agent_id = ? AND customer_id = ? AND chat_status =?", agent_Id, customer_Id, "active").Find(&channel).Error; err != nil {
		return channel, err
	}

	return channel, nil
}

func UpdateChannel(channel models.Channel) (models.Channel, error) {
	if err := config.DB.Save(&channel).Error; err != nil {
		return channel, err
	}

	return channel, nil
}

func GetOneChannelByAgentAndCustomerId(agent_Id, customer_Id int) (models.Channel, error) {
	var channel models.Channel
	if err := config.DB.Where("agent_id = ? AND customer_id = ?", agent_Id, customer_Id).Find(&channel).Error; err != nil {
		return channel, err
	}

	return channel, nil
}
