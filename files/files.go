package files

import (
	"os"
)

// IsFile checks if the given path corresponds to a file.
// It returns true if the path is a file, false if it is a directory, and panics if there is an error.
func IsFile(path string) bool {
	fileInfo, err := os.Stat(path)
	if err != nil {
		panic(err)
	}
	if fileInfo.IsDir() {
		return false
	}
	return true
}
