package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"runtime"
	"strconv"
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

var (
	serverPort        uint16 = 8080 // TODO: Make user-selectable
	lastKnownLocation LocationUpdate
)

func main() {
	defer listen()
}

/* listen spins up a webserver and listens for incoming connections */
func listen() {
	serverIdentifier := fmt.Sprintf("%s on %s %s", runtime.Version(), runtime.GOOS, runtime.GOARCH)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Server", serverIdentifier)
		w.Header().Add("Content-Type", "text/plain")
		w.WriteHeader(501) // Not implemented
		fmt.Fprintf(w, "Sorry, %s is not implemented. Possible options are: /submit to submit a location update or /retrieve to retrieve the last known location", r.RequestURI)
	})

	http.HandleFunc("/retrieve", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Server", serverIdentifier)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(200)
		responseData, err := json.Marshal(&lastKnownLocation)
		if err != nil {
			log.Fatal(err)
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

		locationUpdate := LocationUpdate{
			retrievedLatitude,
			retrievedLongitude,
			retrievedTimestamp,
			retrievedHdop,
			retrievedAltitude,
			retrievedSpeed,
		}

		// Checks if the data (in 'locationUpdate') conforms to the types of the struct 'LocationUpdate'
		_, err := json.Marshal(&locationUpdate)
		if err != nil {
			log.Fatal(err)
		}

		lastKnownLocation = locationUpdate
		location, err := json.Marshal(&lastKnownLocation)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(string(location[:]))
	})

	var listenAddr string = fmt.Sprintf(":%d", serverPort)
	err := http.ListenAndServe(listenAddr, nil)
	if err != nil {
		log.Fatal(err)
	}

}
