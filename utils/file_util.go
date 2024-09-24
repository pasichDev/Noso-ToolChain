package utils

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
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

func CheckIfFileExists(filePath string) bool {
	info, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		return false
	}
	if err != nil {
		// Якщо сталася інша помилка (наприклад, неможливо отримати інформацію про файл)
		log.Printf("Помилка при доступі до файлу: %v", err)
		return false
	}
	// Перевіряємо чи це не папка
	return !info.IsDir()
}

func GetBlockNumberForFile(filePath string) (int, error) {
	fileName := filepath.Base(filePath)

	trimmedName := strings.TrimSuffix(fileName, ".blk")

	return strconv.Atoi(trimmedName)
}
