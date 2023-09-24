package helpers

import (
	"os"
	"path/filepath"
)

// FileList is a function that takes in the pattern and returns an array of matched files over a channel
func FileList(globPattern string) ([]string, error) {
	var files []string

	err := filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && matchesPattern(path, globPattern) {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return files, nil
}

func matchesPattern(path, pattern string) bool {
	match, err := filepath.Match(pattern, filepath.Base(path))
	if err != nil {
		return false
	}
	return match
}
