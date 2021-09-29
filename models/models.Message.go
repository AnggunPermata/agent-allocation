package models

import "gorm.io/gorm"

type Message struct {
	gorm.Model
	ChannelID      uint
	Sender_Role    string `json:"sender_role" form:"sender_role"`
	SenderID       uint   //agent or customer
	Recipient_Role string `json:"recipient_role" form:"recipient_role"`
	RecipientID    uint   //agent or customer
	TextMessage    string `json:"text_message" form:"text_message"`
	Chat_Status    string `json:"chat_status" form:"chat_status"`
}

type Input_Message struct {
	ChannelID      uint
	Sender_Role    string `json:"sender_role" form:"sender_role"`
	SenderID       uint   //agent or customer
	Recipient_Role string `json:"recipient_role" form:"recipient_role"`
	RecipientID    uint   //agent or customer
	TextMessage    string `json:"text_message" form:"text_message"`
}