package initializers

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectToMySQL() (*gorm.DB, error) {
	dsn := "user:password@tcp(127.0.0.1:3306)/database_name?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Printf("Error connecting to MySQL database: %v\n", err)
		return nil, err
	}
	return db, nil
}
