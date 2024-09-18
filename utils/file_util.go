package utils

import (
	"fmt"
	"os"
)

func ReadFile(filePath string) ([]byte, error) {

	if !pathExists(filePath) {
		fmt.Println("Error: File does not exist in the selected path")
		os.Exit(1)
	}
	return os.ReadFile(filePath)
}

func pathExists(path string) bool {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return err == nil
}
