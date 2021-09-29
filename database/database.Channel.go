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
	if err := config.DB.Find(&channelData, "id = ?", channelId).Error; err != nil {
		return channelData, err
	}

	return channelData, nil
}

func GetOneChannelById(channelId int) (models.Channel, error) {
	var channel models.Channel
	if err := config.DB.Find(&channel, "id = ?", channelId).Error; err != nil {
		return channel, err
	}

	return channel, nil
}
