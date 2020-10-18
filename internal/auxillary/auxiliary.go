package auxillary

import (
	"flag"
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
