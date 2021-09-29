package models

import "gorm.io/gorm"

type Customer struct {
	gorm.Model
	Username    string `json:"username" form:"username"`
	Password    string `json:"password" form:"password"`
	Token       string
	Chat_Status string `json:"chat_status" form:"chat_status"` //current chat status of the customer (queue, resolved, active)
}
