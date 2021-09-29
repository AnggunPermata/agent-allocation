package database

import (
	"github.com/anggunpermata/agent-allocation/auth"
	"github.com/anggunpermata/agent-allocation/config"
	"github.com/anggunpermata/agent-allocation/models"
	"github.com/labstack/echo"
)

//AgentLogin used to check agent data in DB, and if exist it will call CreateToken function to initiate new token.
func AgentLogin(username, password string) (models.Agent, error) {
	var err error
	var agent models.Agent
	if err = config.DB.Where("username=? AND password=?", username, password).First(&agent).Error; err != nil {
		return agent, err
	}

	agent.Token, err = auth.CreateToken(int(agent.ID))

	if err != nil {
		return agent, err
	}
	if err := config.DB.Save(agent).Error; err != nil {
		return agent, err
	}
	return agent, nil
}

func GetOneAgent(c echo.Context) (models.Agent, error) {
	var agent models.Agent
	if err := config.DB.Order("count_active_channel asc").Where("agent_status = ?", "active").Find(&agent).Error; err != nil {
		newAgent, err2 := GetAllAgent(c)
		if err2 != nil {
			return newAgent, err2
		}
	}
	return agent, nil
}

func GetAllAgent(c echo.Context) (models.Agent, error) {
	var agent models.Agent
	if err := config.DB.Order("count_active_channel asc").Limit(1).Find(&agent).Error; err != nil {
		return agent, err
	}
	return agent, nil
}

func UpdateAgent(agent models.Agent) (models.Agent, error) {
	if err := config.DB.Save(&agent).Error; err != nil {
		return agent, err
	}

	return agent, nil
}
