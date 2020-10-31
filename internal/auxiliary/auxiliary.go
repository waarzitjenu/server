// Package auxiliary contains some additional, supportive code for this application.
package auxiliary

import (
	"flag"
)

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
