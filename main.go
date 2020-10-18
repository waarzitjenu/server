package main

import (
	"flag"
	"go-osmand-tracker/internal/auxillary"
	"go-osmand-tracker/internal/database"
	"go-osmand-tracker/internal/server"
	"go-osmand-tracker/internal/settings"
	"log"
)

var (
	serverPort uint
	IsDebug    bool
)

const (
	defaultPort        = 8080
	portArgDescription = "The port used to run the application. Defaults to 8080"
	settingsFile       = "settings.json"
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

	// if there is no flags, depend on config file instead of flags
	if !auxillary.IsFlagPassed("p") && !auxillary.IsFlagPassed("port") {

		configFile, err := settings.Read(settingsFile)

		if err != nil {
			log.Printf("Error in settings file %v, settings file will be created by default values!\n", err)
		}

		if err == nil {
			serverPort = configFile.Port
			IsDebug = configFile.Debug
		}
	}

	// create settings file if not exists or corrupted
	if !auxillary.DoesDirExist(settingsFile) || settings.IsCorrupted(settingsFile) {
		config := settings.Config{
			Port:  serverPort,
			Debug: IsDebug,
		}
		err := settings.Write(settingsFile, &config)
		if err != nil {
			log.Printf("Error writing settings file %v\n", err)
		}
	}

	db, err := database.OpenDB("./database", "locations.db")
	if err != nil {
		log.Fatal(err)
	}

	server.SetEnvironment(IsDebug)
	defer server.Listen(serverPort, db)
}
