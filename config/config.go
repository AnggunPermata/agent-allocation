package config

import (
	"strconv"

	"github.com/anggunpermata/agent-allocation/constant"
	"github.com/anggunpermata/agent-allocation/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

var PORT int

func InitDB() {
	// connectionString := "root:12345@tcp(172.17.0.1:3307)/qiscus?charset=utf8&parseTime=True&loc=Local"
	// fmt.Println("constant.Configuration: ", constant.Configuration)
	// fmt.Println(connectionString)
	var err error
	DB, err = gorm.Open(mysql.Open("root:12345@tcp(172.17.0.1:3307)/qiscus?charset=utf8&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	InitMigrate()
}

func InitPort() {
	PORT, _ = strconv.Atoi(constant.Configuration["PORT"])
}

func InitMigrate() {
	DB.AutoMigrate(&models.Customer{})
	DB.AutoMigrate(&models.Agent{})
	DB.AutoMigrate(&models.Channel{})
	DB.AutoMigrate(&models.Message{})
}
