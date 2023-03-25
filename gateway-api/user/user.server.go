package user

import (
	"aqrus/Microservice/gateway-api/utils"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RunServer(db *gorm.DB) {
	config, err := utils.LoadConfig(".")
	if err != nil {
		fmt.Printf("error, %v", err)
		log.Fatal("False to get env", err)
	}

	repo := NewUserRepository(db)
	service := NewUserService(repo)
	controller := NewUserController(service)

	router := gin.Default()
	router.GET("/users", controller.GetAll)
	router.GET("/users/:id", controller.GetByID)
	router.POST("/users", controller.Create)
	router.PUT("/users", controller.Update)

	router.Run(config.ServerAddress)
}
