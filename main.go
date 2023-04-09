package main

import (
	"go-delivery/app/router"
	"go-delivery/config"
)

// @title           go-delivery
// @version         1.0
// @description     go-delivery

// @host      localhost:8080
// @BasePath  /api/v1

// @securityDefinitions.basic  BasicAuth
func main() {
	config.LoadConfig()

	config.InitDB()

	r := router.SetupRouter()

	_ = r.Run(":8080")
}
