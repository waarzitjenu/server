package database

import (
	"errors"
	"go-osmand-tracker/internal/filesystem"
	"strings"

	"github.com/asdine/storm"
)

const (
	defaultDirectory string = "./database"
	defaultFilename  string = "locations.db"
)

// OpenDB opens a new database, or creates one when it doesn't exist. (i.e. a new .db file will be created)
func OpenDB(directory string, filename string) (database *storm.DB, err error) {
	if len(directory) == 0 {
		directory = defaultDirectory
	}

	if len(filename) == 0 {
		filename = defaultFilename
	}

	dirErr := filesystem.CreateDirIfNotExist(directory)
	if dirErr != nil {
		return nil, dirErr
	}

	location := strings.TrimRight(directory, "/") + "/" + filename
	db, dbErr := storm.Open(location)
	if dbErr != nil {
		return nil, dbErr
	}
	return db, nil
}

// DestroyDB deletes an existing database.
func DestroyDB() error {
	return errors.New("Function DestroyDB not implemented")
}
