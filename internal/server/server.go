package server

import (
	"encoding/json"
	"fmt"
	"go-osmand-tracker/internal/types"
	"log"
	"net/http"
	"runtime"
	"strconv"

	"github.com/asdine/storm"
)

var (
	serverIdentifier string = fmt.Sprintf("%s on %s %s", runtime.Version(), runtime.GOOS, runtime.GOARCH)
)

// Listen spins up a webserver and listens for incoming connections
func Listen(port uint, db *storm.DB) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Server", serverIdentifier)
		w.Header().Add("Content-Type", "text/plain")
		w.WriteHeader(501) // Not implemented
		fmt.Fprintf(w, "Sorry, %s is not implemented.", r.RequestURI)
	})

	http.HandleFunc("/retrieve", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Server", serverIdentifier)
		w.Header().Add("Content-Type", "application/json")

		var entry []types.Entry
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

		var entries []types.Entry

		db.All(&entries, storm.Limit(int(cnt)), storm.Reverse())
		responseData, err := json.Marshal(entries)
		if err != nil {
			log.Fatal("Processing entries from database failed", err)
		}
		w.Write(responseData)
	})

	http.HandleFunc("/submit", func(w http.ResponseWriter, r *http.Request) {
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

		locationUpdate := types.LocationUpdate{
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
		entry := types.Entry{
			Timestamp: locationUpdate.Timestamp,
			Data:      locationUpdate,
		}
		err = db.Save(&entry)
		if err != nil {
			log.Fatal("Saving entry to database failed ", err)
		}
	})

	var listenAddr string = fmt.Sprintf(":%d", port)
	fmt.Printf("Starting server at port: %v\n", port)

	err := http.ListenAndServe(listenAddr, nil)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()
}
