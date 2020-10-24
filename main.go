package main

import (
	"flag"
	"go-osmand-tracker/internal/auxiliary"
	"go-osmand-tracker/internal/database"
	"go-osmand-tracker/internal/filesystem"
	"go-osmand-tracker/internal/server"
	"go-osmand-tracker/internal/settings"
	"go-osmand-tracker/internal/types"
	"log"
)

var (
	// IsDebug tells if the server is running in debug mode, i.e. whether or not to provide output messages.
	IsDebug      bool // TODO: Move exported var to main.go
	serverPort   uint
	settingsFile string
)

const (
	settingFileDescription = "The settings file used to run the application. with configuration port and etc."
	defaultSettingsFile    = "settings.json"
	defaultServerPort      = 8080
	defaultDebugStatus     = false
)

func main() {
	// Let the user pick the settings file (optional)
	flag.StringVar(&settingsFile, "config", defaultSettingsFile, settingFileDescription)
	flag.Parse()

	// Create settings file if config not passed, not exists or corrupted
	if !auxiliary.IsFlagPassed("config") && (!filesystem.DoesDirExist(settingsFile) || settings.IsCorrupted(settingsFile)) {
		config := types.Config{
			Port:  defaultServerPort,
			Debug: defaultDebugStatus,
		}
		log.Printf("Initialising settings file %s\n", settingsFile)
		err := settings.Write(settingsFile, &config)
		if err != nil {
			log.Printf("Error writing settings file: %s\n", err)
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
