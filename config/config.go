package config

import (
	"log"
	"os"
	"strconv"

	"github.com/anggunpermata/agent-allocation/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

var PORT int

func goDotEnvVariable(key string) string {

	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

func InitDB() {
	// connectionString := "root:12345@tcp(172.17.0.1:3307)/qiscus?charset=utf8&parseTime=True&loc=Local"
	// fmt.Println("constant.Configuration: ", constant.Configuration)
	// fmt.Println(connectionString)
	var err error
	DB, err = gorm.Open(mysql.Open(goDotEnvVariable("DATABASE_URL")), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	InitMigrate()
}

func InitPort() {
	PORT, _ = strconv.Atoi(goDotEnvVariable("PORT"))
}

func InitMigrate() {
	DB.AutoMigrate(&models.Customer{})
	DB.AutoMigrate(&models.Agent{})
	DB.AutoMigrate(&models.Channel{})
	DB.AutoMigrate(&models.Message{})
}
