package controller

import (
	"net/http"

	"github.com/anggunpermata/agent-allocation/auth"
	"github.com/anggunpermata/agent-allocation/database"
	"github.com/anggunpermata/agent-allocation/models"
	"github.com/labstack/echo"
)

//To check user's authorization by using user Id

func AuthorizedCustomer(c echo.Context) bool {
	_, role := auth.ExtractTokenUserId(c)
	if role != "agent" {
		return false
	}
	return true
}

func CustomerLogin(c echo.Context) error {
	inputData := models.Login_Form{}
	c.Bind(&inputData)
	userData := models.Customer{
		Username: inputData.Username,
		Password: inputData.Password,
	}
	c.Bind(&userData)

	customer, err := database.CustomerLogin(userData.Username, userData.Password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "please check your username and password again.",
		})
	}

	showCustomerData := map[string]interface{}{
		"ID":       customer.ID,
		"Username": "@" + customer.Username,
		"Token":    customer.Token,
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":   "Hello, Let's Start!",
		"user-data": showCustomerData,
	})
}
