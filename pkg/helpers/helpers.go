package helpers

import (
	"path/filepath"

	log "github.com/sirupsen/logrus"
)

// FileList is a function that takes in the pattern and returns an array of matched files over a channel
func FileList(globPattern string, logLevel log.Level) ([]string, error) {
	files, err := filepath.Glob(globPattern)
	if err != nil {
		log.Error("error processing the pattern:", err)
	}
	return files, err
}
