package util

import (
	"os"
)

func BuildPath(relative_path string) (string, error) {
	path, err := os.Getwd()
	return path, err
}

func GetResource(filename string) (string, error) {
	return BuildPath("/res/" + filename)
}