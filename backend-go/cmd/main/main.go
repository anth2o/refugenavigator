package main

import (
	"log"

	"github.com/anth2o/refugenavigator/internal/server"
)

func main() {
	engine := server.SetupRoutes()
	if err := engine.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
