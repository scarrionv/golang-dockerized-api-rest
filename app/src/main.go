package main

import (
	"github.com/scarrionv/simple-api/app/src/controllers"
	log "github.com/sirupsen/logrus"
	"os"
)

const (
	EnvPort     = "PORT"
	defaultPort = "8888"
)

func main() {
	var serverPort = defaultPort
	var port, present = os.LookupEnv(EnvPort)
	if present {
		serverPort = port
	}
	log.Info("[main] Starting app on port: " + serverPort)
	controllers.CreateControllers(serverPort)
}
