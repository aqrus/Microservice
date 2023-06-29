package user

import (
	"aqrus/Microservice/gateway-api/utils"
	"fmt"
	"log"
	"golang.org/x/time/rate"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RunServer(db *gorm.DB) {
	config, err := utils.LoadConfig(".")
	if err != nil {
		fmt.Printf("error, %v", err)
		log.Fatal("False to get env", err)
	}
	// limiter := rate.NewLimiter(rate.Limit(config.RateLimit), config.RateLimitBurst)

	repo := NewUserRepository(db)
	service := NewUserService(repo)
	controller := NewUserController(service)

	router := gin.Default()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(func(c *gin.Context) {
		// c.Set("redisClient", redisClient)
		// c.Set("limiter", limiter)
		c.Next()
	})
	router.GET("/users", controller.GetAll)
	router.GET("/users/:id", controller.GetByID)
	router.POST("/users", controller.Create)
	router.PUT("/users", controller.Update)

	router.Run(config.ServerAddress)
}
