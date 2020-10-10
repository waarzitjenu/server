package main

import (
	"flag"
	"go-osmand-tracker/internal/database"
	"go-osmand-tracker/internal/server"
	"log"
)

var (
	serverPort uint
	IsDebug    bool
)

const (
	defaultPort        = 8080
	portArgDescription = "The port used to run the application. Defaults to 8080"
)

func main() {
	// let the user pick the port by using "port" or "p" option
	flag.UintVar(&serverPort, "port", defaultPort, portArgDescription)
	flag.UintVar(&serverPort, "p", defaultPort, portArgDescription+" (shorthand)")
	debugMode := flag.Bool("debug", false, "Log messages to stdout")
	flag.Parse()

	if *debugMode {
		IsDebug = true
	}

	db, err := database.OpenDB("./database", "locations.db")
	if err != nil {
		if IsDebug {
			log.Fatal(err)
		}
	}

	server.SetEnvironment(IsDebug)
	defer server.Listen(serverPort, db)
}
