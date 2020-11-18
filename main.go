package main

import (
	"encoding/json"
	"flag"
	"fmt"

	"log"
	"net/http"
	"runtime"
	"strconv"

	"crypto/tls"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/waarzitjenu/database"
	"github.com/waarzitjenu/database/filesystem"
	"github.com/waarzitjenu/database/types"
	"github.com/waarzitjenu/server/internal/settings"
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
	serverIdentifier string = fmt.Sprintf("%s on %s %s", runtime.Version(), runtime.GOOS, runtime.GOARCH)
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

	err = core.OpenDatabase("./database", "locations.db")
	if err != nil {
		log.Fatal(err)
	}

	defer Listen(configFile)
}

// IsFlagPassed checks if a given flag is used via the CLI and returns a boolean with the value true when the flag has been passed.
func IsFlagPassed(name string) bool {
	// falg.Visit shall be called after flags is parsed
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

// Listen spins up a webserver and listens for incoming connections
func Listen(config *types.Config) {

	if config.Debug {
		gin.SetMode(gin.DebugMode)
		core.Debug = true
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	ginEngine := gin.New()        // Creates a router without any middleware by default
	ginEngine.Use(gin.Recovery()) // Recovery middleware recovers from any panics and writes a 500 if there was one.
	if config.Debug {
		ginEngine.Use(gin.Logger())
	} // Only print access logs when debug mode is active.

	if filesystem.DoesDirExist("./web/dist") {
		ginEngine.Use(static.Serve("/", static.LocalFile("./web/dist", false)))
	} else {
		ginEngine.GET("/", func(c *gin.Context) {
			c.Header("Server", serverIdentifier)
			c.Header("Content-Type", "text/plain")
			c.Writer.WriteHeader(501)
			c.String(http.StatusNotImplemented, "Sorry, %s is not implemented.", c.Request.RequestURI)
		})
	}

	ginEngine.GET("/retrieve", func(c *gin.Context) {
		c.Header("Server", serverIdentifier)
		c.Header("Content-Type", "application/json")
		c.Header("Access-Control-Allow-Origin", "*")
		c.Writer.WriteHeader(200)

		var (
			entries         []types.DatabaseEntry
			numberOfEntries uint16 = 1
		)

		c.Request.URL.Query().Get("count")
		if len(c.Request.URL.Query().Get("count")) >= 1 {
			parsedCountValue, err := strconv.ParseUint(c.Request.URL.Query().Get("count"), 10, 16)
			if err != nil {
				if config.Debug {
					log.Println("[ERROR] Error parsing query parameter 'count'", err)
				}
			} else {
				numberOfEntries = uint16(parsedCountValue)
			}
		}

		// Let's fetch the entries from the database
		entries, err := core.RetrieveEntries(int(numberOfEntries))
		if err != nil {
			log.Fatal(err)
		}

		// Perform type checking
		_, jsonErr := json.Marshal(entries)
		if jsonErr != nil {
			log.Fatal(jsonErr)
		}

		// Write a response body
		NegotiatedResponseBody(c, entries)
	})

	ginEngine.GET("/submit", func(c *gin.Context) {
		c.Header("Server", serverIdentifier)
		c.Writer.WriteHeader(204) // The server successfully processed the request, and is not returning any content.

		retrievedLatitude, _ := strconv.ParseFloat(c.Request.URL.Query().Get("lat"), 64)
		retrievedLongitude, _ := strconv.ParseFloat(c.Request.URL.Query().Get("lon"), 64)
		retrievedTimestamp, _ := strconv.ParseUint(c.Request.URL.Query().Get("timestamp"), 10, 64)
		retrievedHdop, _ := strconv.ParseFloat(c.Request.URL.Query().Get("hdop"), 64)
		retrievedAltitude, _ := strconv.ParseFloat(c.Request.URL.Query().Get("altitude"), 64)
		retrievedSpeed, _ := strconv.ParseFloat(c.Request.URL.Query().Get("speed"), 64)

		err := core.SaveEntry(types.LocationUpdate{
			Latitude:  retrievedLatitude,
			Longitude: retrievedLongitude,
			Timestamp: retrievedTimestamp,
			Hdop:      retrievedHdop,
			Altitude:  retrievedAltitude,
			Speed:     retrievedSpeed,
		})
		if err != nil {
			log.Println(err)
		}
	})

	var listenAddr string = fmt.Sprintf(":%d", config.ServerConfiguration.Port)
	if config.Debug {
		log.Println("[WARNING] Debug mode is enabled. Disable it in production environments to prevent logging sensitive data.")
		log.Printf("[INFO] Starting server on port %v\n", config.ServerConfiguration.Port)
	}

	var err error
	server := http.Server{
		Addr:    listenAddr,
		Handler: ginEngine,
	}

	if config.ServerConfiguration.TLS.Enabled {
		tlsConfig := &tls.Config{
			MinVersion:       tls.VersionTLS12,
			CurvePreferences: []tls.CurveID{tls.CurveP521, tls.CurveP384, tls.CurveP256},
		}

		server.TLSConfig = tlsConfig

		// validate if the crt & key is added
		if len(config.ServerConfiguration.TLS.Certificate.PublicKey) == 0 || len(config.ServerConfiguration.TLS.Certificate.PublicKey) == 0 {
			log.Fatal("[FATAL] Invalid certificate. Please define the TLS certificate and key in settings.json!")
		}

		err = server.ListenAndServeTLS(config.ServerConfiguration.TLS.Certificate.PublicKey, config.ServerConfiguration.TLS.Certificate.PrivateKey)
	} else {
		if config.Debug {
			log.Println(`[WARNING] Please note that the server is running without SSL/TLS security. It is highly recommended to enable it in production environments. Refer to the project documentation to learn how to set up TLS.`)
		}
		err = server.ListenAndServe()
	}

	if err != nil {
		log.Fatal(err)
	}
}

// NegotiatedResponseBody checks the value of the 'Accept' header and responds with the requested format, if supported (JSON, XML or YAML). Defaults to JSON.
func NegotiatedResponseBody(c *gin.Context, entries []types.DatabaseEntry) {
	switch c.NegotiateFormat(gin.MIMEJSON, gin.MIMEXML, gin.MIMEYAML) {
	case gin.MIMEXML:
		c.XML(http.StatusOK, entries)
	case gin.MIMEYAML:
		c.YAML(http.StatusOK, entries)
	case gin.MIMEJSON:
		c.JSON(http.StatusOK, entries)
	}
}
