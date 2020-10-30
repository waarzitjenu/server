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

const (
	settingFileDescription = "The settings file used to run the application. with configuration port and etc."
	defaultSettingsFile    = "settings.json"
)

var (
	settingsFile         string
	defaultConfiguration types.Config = types.Config{
		Debug: true,
		ServerConfiguration: types.ServerConfiguration{
			Port: 8080,
			TLS: types.TLS{
				Enabled: false,
			},
		},
	}
)

func main() {
	// Let the user pick the settings file (optional)
	flag.StringVar(&settingsFile, "config", defaultSettingsFile, settingFileDescription)
	flag.Parse()

	// Create settings file if config not passed, not exists or corrupted
	if !auxiliary.IsFlagPassed("config") && (!filesystem.DoesDirExist(settingsFile) || settings.IsCorrupted(settingsFile)) {
		log.Printf("Initialising settings file %s\n", settingsFile)
		config := defaultConfiguration
		err := settings.Write(settingsFile, &config)
		if err != nil {
			log.Printf("Error writing settings file: %s\n", err)
		}
	}

	configFile, err := settings.Read(settingsFile)
	if err != nil {
		log.Fatal(err)
	}

	db, err := database.OpenDB("./database", "locations.db")
	if err != nil {
		log.Fatal(err)
	}

	defer server.Listen(configFile, db)
}
