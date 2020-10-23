// Package auxiliary contains some additional, supportive code for this application.
package auxiliary

import (
	"flag"
)

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
