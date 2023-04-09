package main

import (
	"go-delivery/app/router"
	"go-delivery/config"
)

func main() {
	config.LoadConfig()

	config.InitDB()

	r := router.SetupRouter()

	_ = r.Run(":8080")
}
