package controller

import (
	"net/http"
	"strconv"

	"github.com/anggunpermata/agent-allocation/auth"
	"github.com/anggunpermata/agent-allocation/database"
	"github.com/anggunpermata/agent-allocation/models"
	"github.com/labstack/echo"
)

//To check user's authorization by using user Id

func AuthorizedCustomer(customerId int, c echo.Context) error {
	// _, role := auth.ExtractTokenUserId(c)
	// if role != "agent" {
	// 	return false
	// }
	// return true

	customer, err := database.GetOneCustomerById(customerId)
	loggedInCustomerId, role := auth.ExtractTokenUserId(c)
	if loggedInCustomerId != int(customer.ID) || err != nil || role != "customer" {
		return echo.NewHTTPError(http.StatusUnauthorized, "Cannot access this account")
	}
	return nil
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

func CustomerLogout(c echo.Context) error {
	customerId, err := strconv.Atoi(c.Param("customer_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid id",
		})
	}
	if err = AuthorizedCustomer(customerId, c); err != nil {
		return err
	}
	logout, err := database.GetOneCustomerById(customerId)
	logout.Token = ""

	c.Bind(&logout)
	customer, err := database.UpdateCustomer(logout)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "cannot logout",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":     "Thank you for using Our service~",
		"Customer ID": customer.ID,
		"Username":    customer.Username,
	})
}
