// Package filesystem contains some general filesystem functions.
package filesystem

import (
	"os"
)

// DoesDirExist checks if a directory exists and returns a boolean.
func DoesDirExist(directory string) bool {
	_, err := os.Stat(directory)
	if err != nil {
		return false
	}
	return true
}

// CreateDirIfNotExist first checks if a directory exists and creates it if does not exist
func CreateDirIfNotExist(dir string) error {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0755)
		if err != nil {
			return err
		}
	}
	return nil
}

// DeleteFile deletes a single file. Backup indicates to rename the file instead of actually deleting it.
func DeleteFile(target string, backup bool) error {
	if backup {
		return os.Rename(target, target+".bak")
	}
	return os.Remove(target)
}

// DeleteDirectory deletes an entire directory.
func DeleteDirectory(target string) error {
	err := os.RemoveAll(target)
	if err != nil {
		return err
	}
	return nil
}

// DeleteDirectoryContents deletes the contents of a directory, but keeps the directory itself.
func DeleteDirectoryContents(target string) error {
	// Open the directory and obtain a list of all of its files.
	dir, err := os.Open(target)
	if err != nil {
		return err
	}
	files, err := dir.Readdir(0)
	if err != nil {
		return err
	}

	// Loop over the directory's files.
	for index := range files {
		file := files[index]

		// Get name of file and its full path.
		filename := file.Name()
		fullPath := target + filename

		// Remove the file.
		err := os.Remove(fullPath)
		if err != nil {
			return err
		}
	}

	return nil
}
