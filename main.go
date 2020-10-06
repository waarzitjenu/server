package main

import (
	"flag"
	"go-osmand-tracker/internal/database"
	"go-osmand-tracker/internal/server"
	"log"
)

var (
	serverPort uint
)

const (
	defaultPort        = 8080
	portArgDescription = "The port used to run the application. Defaults to 8080"
)

func main() {
	// let the user pick the port by using "port" or "p" option
	flag.UintVar(&serverPort, "port", defaultPort, portArgDescription)
	flag.UintVar(&serverPort, "p", defaultPort, portArgDescription+" (shorthand)")
	flag.Parse()

	db, err := database.OpenDB("./database", "locations.db")
	if err != nil {
		log.Fatal(err)
	}
	defer server.Listen(serverPort, db)
}
