// Package engine handles the public API endpoints and their linked functionalities.
package engine

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"runtime"
	"strconv"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/waarzitjenu/server/internal/database"
	"github.com/waarzitjenu/server/internal/filesystem"
	"github.com/waarzitjenu/server/internal/types"
)

var (
	// ServerIdentifier is the value of the 'Server' header sent from the server to the clients.
	ServerIdentifier string = fmt.Sprintf("%s on %s %s", runtime.Version(), runtime.GOOS, runtime.GOARCH)
)

// Listen spins up a webserver and listens for incoming connections
func Listen(config *types.Config) {

	if config.Debug {
		gin.SetMode(gin.DebugMode)
		database.Debug = true
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
			c.Header("Server", ServerIdentifier)
			c.Header("Content-Type", "text/plain")
			c.Writer.WriteHeader(501)
			c.String(http.StatusNotImplemented, "Sorry, %s is not implemented.", c.Request.RequestURI)
		})
	}

	ginEngine.GET("/retrieve", func(c *gin.Context) {
		c.Header("Server", ServerIdentifier)
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
		entries, err := database.Read(int(numberOfEntries))
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
		c.Header("Server", ServerIdentifier)
		c.Writer.WriteHeader(204) // The server successfully processed the request, and is not returning any content.

		retrievedLatitude, _ := strconv.ParseFloat(c.Request.URL.Query().Get("lat"), 64)
		retrievedLongitude, _ := strconv.ParseFloat(c.Request.URL.Query().Get("lon"), 64)
		retrievedTimestamp, _ := strconv.ParseUint(c.Request.URL.Query().Get("timestamp"), 10, 64)
		retrievedHdop, _ := strconv.ParseFloat(c.Request.URL.Query().Get("hdop"), 64)
		retrievedAltitude, _ := strconv.ParseFloat(c.Request.URL.Query().Get("altitude"), 64)
		retrievedSpeed, _ := strconv.ParseFloat(c.Request.URL.Query().Get("speed"), 64)

		err := database.Create(types.LocationUpdate{
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
