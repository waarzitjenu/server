package main

import (
	"flag"

	"log"

	"github.com/waarzitjenu/server/internal/database"
	"github.com/waarzitjenu/server/internal/engine"
	"github.com/waarzitjenu/server/internal/settings"
	"github.com/waarzitjenu/server/internal/types"
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
	if !IsFlagPassed("config") && settings.IsCorrupted(settingsFile) {
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

	// First, we're going to open a database. If that is successful, we will receive a pointer.
	db, err := database.Open("./database/locations.db")
	if err != nil {
		log.Fatal(err)
	}

	// Now, we must tell the database module to use this database.
	err = database.Use(db)
	if err != nil {
		log.Fatal(err)
	}

	defer engine.Listen(configFile)
}

// IsFlagPassed checks if a given flag is used via the CLI and returns a boolean with the value true when the flag has been passed.
func IsFlagPassed(name string) bool {
	// flag.Visit shall be called after flags is parsed
	if !flag.Parsed() {
		flag.Parse()
	}

	found := false
	flag.Visit(func(f *flag.Flag) {
		if f.Name == name {
			found = true
		}
	})

	return found
}
