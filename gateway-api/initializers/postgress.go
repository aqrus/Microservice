package initializers

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectToPostgres() (*gorm.DB, error) {
	dsn := "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Failed to connect to Postgres database!")
		return nil, err
	}
	fmt.Println("Connected to Postgres database!")
	return db, nil
}
