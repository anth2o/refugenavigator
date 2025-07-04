package main

import (
	"log"

	"github.com/anth2o/refugenavigator/internal/server"
)

func main() {
	engine := server.SetupRoutes()
	if err := engine.Run(); err != nil {
		log.Fatal(err)
	}
}
