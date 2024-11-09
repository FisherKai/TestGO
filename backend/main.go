package main

import (
	"demo_hubu_backend/api"
	"demo_hubu_backend/config"
	"demo_hubu_backend/utils"
	"fmt"
)

func main() {
	db, err := config.DbInit()
	if err != nil {
		fmt.Println("connection database error")
	}
	r := api.SetupRouter(db)
	r.Use(utils.LoggerMiddleware())
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
