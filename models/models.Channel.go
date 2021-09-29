package models

import "gorm.io/gorm"

type Channel struct {
	gorm.Model
	CustomerID  uint
	AgentID     uint
	Chat_Status string `json:"chat_status" form:"chat_status"`
}
