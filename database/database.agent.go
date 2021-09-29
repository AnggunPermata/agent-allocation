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
	return agent, nil
}

func GetOneAgent(c echo.Context) (models.Agent, error) {
	var agent models.Agent
	//first check through active agent
	if err := config.DB.Find(&agent, "agent_status=?", "active").Error; err != nil {
		//if there's no active agent, get the minimum agent
		config.DB.Raw("select min(count_active_channel) from agent").Scan(&agent)
		return agent, nil
	}
	//if there exist active agent, allocate the agent to channel
	config.DB.Raw("select min(count_active_channel) from agent where agent_status = ?", "active").Scan(&agent)
	return agent, nil

}
