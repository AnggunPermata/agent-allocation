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
