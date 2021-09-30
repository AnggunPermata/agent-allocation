package database

import (
	"github.com/anggunpermata/agent-allocation/config"
	"github.com/anggunpermata/agent-allocation/models"
)

func CreateNewMessage(message models.Message) (models.Message, error) {
	if err := config.DB.Save(&message).Error; err != nil {
		return message, err
	}
	return message, nil
}

func GetLastMessageByChannelId(channelId int) (models.Message, error) {
	var message models.Message
	if err := config.DB.Order("created_at desc").Where("channel_id = ?", channelId).Find(&message).Limit(1).Error; err != nil {
		return message, err
	}
	return message, nil
}

func UpdateChatStatus(message models.Message) (models.Message, error) {
	if err := config.DB.Save(&message).Error; err != nil {
		return message, err
	}
	return message, nil
}
