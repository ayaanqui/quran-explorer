package util

import (
	"os"
)

func BuildPath(relative_path string) (string, error) {
	path, err := os.Getwd()
	full_path := path + relative_path
	return full_path, err
}

// Returns full path for a given filename.
// Assumes the file destination to be in the res folder of the project root
func GetResource(filename string) (string, error) {
	return BuildPath("/res/" + filename)
}