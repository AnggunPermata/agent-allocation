package database

import (
	"github.com/anggunpermata/agent-allocation/auth"
	"github.com/anggunpermata/agent-allocation/config"
	"github.com/anggunpermata/agent-allocation/models"
)

//CustomerLogin used to check user data in DB, and if exist it will call CreateToken function to initiate new token.
func CustomerLogin(username, password string) (models.Customer, error) {
	var err error
	var customer models.Customer
	if err = config.DB.Where("username=? AND password=?", username, password).First(&customer).Error; err != nil {
		return customer, err
	}

	customer.Token, err = auth.CreateCustomerToken(int(customer.ID))

	if err != nil {
		return customer, err
	}
	if err := config.DB.Save(customer).Error; err != nil {
		return customer, err
	}
	return customer, nil
}

func GetOneCustomerById(customerId int) (models.Customer, error) {
	var customer models.Customer
	if err := config.DB.Where("id=?", customerId).First(&customer).Error; err != nil {
		return customer, err
	}
	return customer, nil
}

func UpdateCustomer(customer models.Customer) (models.Customer, error) {
	if err := config.DB.Save(&customer).Error; err != nil {
		return customer, err
	}

	return customer, nil
}
