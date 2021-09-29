package models

import "gorm.io/gorm"

type Agent struct {
	gorm.Model
	Username             string `json:"username" form:"username"`
	Password             string `json:"password" form:"password"`
	Count_Active_Channel int    //how many channels being served by the agent at the time
	Agent_Status         string `json:"agent_status" form:"agent_status"` //is agent active or not
	Token                string
}

type Login_Form struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}
