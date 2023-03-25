package initializers

import (
	"fmt"
	"log"

	"aqrus/Microservice/gateway-api/utils"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToBD() {
	config, err := utils.LoadConfig(".")
	if err != nil {
		fmt.Printf("error, %v", err)
		log.Fatal("False to get env", err)
	}

	dsn := fmt.Sprintf("host=localhost user=%s password=%s dbname=%s port=5432 sslmode=disable", config.PostgresUser, config.PostgresPassword, config.DBName)
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("False to connect to database")
	}
}
