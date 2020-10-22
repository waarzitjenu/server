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
	serverPort   uint
	IsDebug      bool
	settingsFile string
)

const (
	settingFileDescription = "The settings file used to run the application. with configuration port and etc."
	defaultSettingsFile    = "settings.json"
	defaultServerPort      = 8080
	defaultDebugStatus     = false
)

func main() {
	// let the user pick the settings file (optional)
	flag.StringVar(&settingsFile, "config", defaultSettingsFile, settingFileDescription)
	flag.Parse()

	// create settings file if config not passed, not exists or corrupted
	if !auxillary.IsFlagPassed("config") && (!auxillary.DoesDirExist(settingsFile) || settings.IsCorrupted(settingsFile)) {
		config := settings.Config{
			Port:  defaultServerPort,
			Debug: defaultDebugStatus,
		}
		err := settings.Write(settingsFile, &config)
		if err != nil {
			log.Printf("Error writing settings file %s: %s\n", settingsFile, err)
		}
	}

	configFile, err := settings.Read(settingsFile)
	if err != nil {
		log.Fatal(err)
	}

	serverPort = configFile.Port
	IsDebug = configFile.Debug

	db, err := database.OpenDB("./database", "locations.db")
	if err != nil {
		log.Fatal(err)
	}

	server.SetEnvironment(IsDebug)
	defer server.Listen(serverPort, db)
}
