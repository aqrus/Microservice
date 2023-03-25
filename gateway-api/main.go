package main

import (
	"aqrus/Microservice/gateway-api/initializers"
	// service "aqrus/Microservice/gateway-api/services"
	"aqrus/Microservice/gateway-api/user"
)

func main() {
	initializers.ConnectToBD()
	// service.RunGRPCServer()
	initializers.DB.AutoMigrate(&user.User{})
	user.RunServer(initializers.DB)
	
}
