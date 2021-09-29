package models

type Customer struct {
	ID          uint   `json:"id" form:"id" gorm:"AUTO_INCREMENT"`
	Username    string `json:"username" form:"username"`
	Password    string `json:"password" form:"password"`
	Token       string
	Chat_Status string `json:"chat_status" form:"chat_status"` //current chat status of the customer (queue, resolved, active)
}
