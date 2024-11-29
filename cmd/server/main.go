package main

import (
	"ecommerce-app/config"
	"ecommerce-app/internal/router"
	"log"
)

func main() {
	config.LoadConfig()
	r := router.SetupRouter()
	log.Fatal(r.Listen(":8000"))
}
