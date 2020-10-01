package main

import (
	"encoding/base64"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime"
	"strconv"
	"strings"

	"github.com/asdine/storm"
)

// LocationUpdate contains the location update data types as retrieved from the OsmAnd app by default
type LocationUpdate struct {
	Latitude  float64 `json:"latitude,omitempty"`
	Longitude float64 `json:"longitude,omitempty"`
	Timestamp uint64  `json:"timestamp,omitempty"`
	Hdop      float64 `json:"hdop,omitempty"`
	Altitude  float64 `json:"altitude,omitempty"`
	Speed     float64 `json:"speed,omitempty"`
}

// Entry contains the database entry. The ID field is automatically incremented and
// Timestamp is indexed to allow fast queries based on time ranges
type Entry struct {
	ID        int    `storm:"id,increment"`
	Timestamp uint64 `storm:"index"`
	Data      LocationUpdate
}

var (
	serverPort   uint
	databasePath string = "./database"
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

	createDirIfNotExist(databasePath)
  
	db, err := storm.Open(databasePath + "/locations.db")
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()
	defer listen(db)
}

/* listen spins up a webserver and listens for incoming connections */
func listen(db *storm.DB) {
	serverIdentifier := fmt.Sprintf("%s on %s %s", runtime.Version(), runtime.GOOS, runtime.GOARCH)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Server", serverIdentifier)
		w.Header().Add("Content-Type", "text/plain")
		w.WriteHeader(501) // Not implemented
		fmt.Fprintf(w, "Sorry, %s is not implemented. Possible options are: /submit to submit a location update or /retrieve to retrieve the last known location", r.RequestURI)
	})

	http.HandleFunc("/retrieve", basicAuth(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Server", serverIdentifier)
		w.Header().Add("Content-Type", "application/json")

		var entry []Entry
		db.All(&entry, storm.Limit(1), storm.Reverse())

		var responseData []byte

		if len(entry) == 0 {
			// There is no data. Let's return HTTP Status 204: No Content
			w.WriteHeader(204) // The server successfully processed the request, but is not returning any content.
		} else {
			// There is data. Let's process it.
			processedEntry, err := json.Marshal(entry[0])
			if err != nil {
				w.WriteHeader(500) // HTTP 500 Internal Server Error
				log.Println("Error retrieving last database entry", err)
			} else {
				w.WriteHeader(200) // HTTP 200 OK
				responseData = processedEntry
			}
		}

		w.Write(responseData)
	})

	http.HandleFunc("/retrieve/multi", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Server", serverIdentifier)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(200)

		var cnt uint16 = 10
		if len(r.URL.Query().Get("count")) > 0 {
			parsedCountValue, err := strconv.ParseUint(r.URL.Query().Get("count"), 10, 16)
			if err != nil {
				log.Println("Error parsing query parameter 'count'", err)
			} else {
				cnt = uint16(parsedCountValue)
			}
		}

		var entries []Entry

		db.All(&entries, storm.Limit(int(cnt)), storm.Reverse())
		responseData, err := json.Marshal(entries)
		if err != nil {
			log.Fatal("Processing entries from database failed", err)
		}
		w.Write(responseData)
	}))

	http.HandleFunc("/submit", basicAuth(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Server", serverIdentifier)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(204) // The server successfully processed the request, and is not returning any content.

		// TODO: Error handling
		retrievedLatitude, _ := strconv.ParseFloat(r.URL.Query().Get("lat"), 64)
		retrievedLongitude, _ := strconv.ParseFloat(r.URL.Query().Get("lon"), 64)
		retrievedTimestamp, _ := strconv.ParseUint(r.URL.Query().Get("timestamp"), 10, 64)
		retrievedHdop, _ := strconv.ParseFloat(r.URL.Query().Get("hdop"), 64)
		retrievedAltitude, _ := strconv.ParseFloat(r.URL.Query().Get("altitude"), 64)
		retrievedSpeed, _ := strconv.ParseFloat(r.URL.Query().Get("speed"), 64)

		locationUpdate := LocationUpdate{
			retrievedLatitude,
			retrievedLongitude,
			retrievedTimestamp,
			retrievedHdop,
			retrievedAltitude,
			retrievedSpeed,
		}

		// Checks if the data (in 'locationUpdate') conforms to the types of the struct 'LocationUpdate'
		_, err := json.Marshal(locationUpdate)
		if err != nil {
			log.Fatal(err)
		}

		// Prepare and insert into DB
		entry := Entry{
			Timestamp: locationUpdate.Timestamp,
			Data:      locationUpdate,
		}
		err = db.Save(&entry)
		if err != nil {
			log.Fatal("Saving entry to database failed", err)
		}
	}))

	var listenAddr string = fmt.Sprintf(":%d", serverPort)
	fmt.Printf("Starting server at port: %v\n", serverPort)

	err := http.ListenAndServe(listenAddr, nil)
	if err != nil {
		log.Fatal(err)
	}
}

// createDirIfNotExist first checks if a directory exists and creates it if does not exist
func createDirIfNotExist(dir string) {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0755)
		if err != nil {
			panic(err)
		}
	}
}

func basicAuth( handler http.HandlerFunc) http.HandlerFunc{

	return func(w http.ResponseWriter, r *http.Request) {

		auth := strings.SplitN(r.Header.Get("Authorization"), " ", 2)

		if len(auth) != 2 || auth[0] != "Basic" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		payload, _ := base64.StdEncoding.DecodeString(auth[1])
		pair := strings.SplitN(string(payload), ":", 2)

		if len(pair) != 2 || !validate(pair[0], pair[1]) {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
	}
}

func validate(username, password string) bool {
	if username == "user" && password == "password" {
		return true
	}
	return false
}