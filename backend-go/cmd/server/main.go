package main

import (
	"github.com/anth2o/refugenavigator/internal/scrapper"
	"github.com/anth2o/refugenavigator/internal/server"
)

func main() {
	connected := scrapper.ConnectDB()
	if connected {
		defer scrapper.CloseDB()
	}
	server.Run()
}
