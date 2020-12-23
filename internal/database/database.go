package database

import (
	"encoding/json"
	"fmt"
	"path"
	"time"

	"github.com/asdine/storm"

	"github.com/waarzitjenu/server/internal/filesystem"
	"github.com/waarzitjenu/server/internal/types"
)

var (
	// Debug tells wheter or not printing debug messages
	Debug bool = false
	// db contains a pointer to the current database in use
	db *storm.DB
)

const (
	defaultDatabaseLocation string = "./database/locations.db"
)

// Open opens an existing database or creates one when it does not already exist. (i.e. a new .db file will be created), eventually, it points global variable `db` to it.
func Open(filename string) (*storm.DB, error) {
	if len(filename) == 0 {
		filename = defaultDatabaseLocation
	}

	dir, _ := path.Split(filename)
	err := filesystem.CreateDirIfNotExist(dir)
	if err != nil {
		err = fmt.Errorf("Error while creating directory %s: %w; ", dir, err)
		return nil, err
	}

	db, err := storm.Open(filename)
	if err != nil {
		err = fmt.Errorf("Error while opening database %s: %w; ", filename, err)
		return nil, err
	}

	return db, nil
}

// Use tells the database module to actually take a BoltDB database into use.
// It does so by pointing global variable `db` to it, so transactions can be made (CRUD).
func Use(database *storm.DB) error {
	if db != nil {
		err := fmt.Errorf("Error while using database %s: Another database is already in use; ", database)
		return err
	}
	db = database
	return nil
}

// Create saves a given entry to the database
func Create(entry types.LocationUpdate) (err error) {
	// Checks if the data (in 'entry') conforms to the types of the struct 'LocationUpdate'
	_, err = json.Marshal(entry)
	if err != nil {
		err = fmt.Errorf("Entry does not conform to the struct LocationUpdate: %w; ", err)
	}

	// Prepare and insert into DB
	dbEntry := types.DatabaseEntry{
		Timestamp: uint64(time.Now().Unix()),
		Data:      entry,
	}

	// Checks if the data (in 'dbEntry') conforms to the types of the struct 'DatabaseEntry'
	_, err = json.Marshal(dbEntry)
	if err != nil {
		err = fmt.Errorf("dbEntry does not conform to the struct DatabaseEntry: %w; ", err)
	}

	// Save the new entry to the database.
	err = db.Save(&dbEntry)
	if err != nil {
		err = fmt.Errorf("Saving the entry to the database failed: %w; ", err)
	}

	// Return an error if any
	return err
}

// CreateMultiple saves multiple entries to the database.
func CreateMultiple(entries ...types.LocationUpdate) error {
	for _, entry := range entries {
		err := Create(entry)
		if err != nil {
			return err
		}
	}
	return nil
}

// Read retrieves entries from the database and returns them
func Read(count int) (entries []types.DatabaseEntry, err error) {
	if Debug {
		fmt.Println("Fetching entries from database")
	}
	db.All(&entries, storm.Limit(int(count)), storm.Reverse())

	_, err = json.Marshal(&entries)
	if err != nil {
		err = fmt.Errorf("Processing entries from database failed: %w; ", err)
	}

	return entries, err
}

// Update updates a single entry of the database. Not implemented yet.
func Update() error {
	// TODO: Write function body.
	return nil
}

// Delete deletes an entry from the database. Not implemented yet.
func Delete() error {
	// TODO: Write function body.
	return nil
}

// Destroy deletes a database from disk
func Destroy(filename string) error {
	return filesystem.DeleteFile(filename, false)
}
