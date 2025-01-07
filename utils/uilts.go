package utils

import (
	"os"
	"path/filepath"
	"strconv"
)

func FileExistsInStorage(filename string) bool {
	storageDir := "./storage/in"
	filePath := filepath.Join(storageDir, filename)
	_, err := os.Stat(filePath)
	return !os.IsNotExist(err)
}

func ToInFilePath(filename string) string {
	storageDir := "./storage/in"
	return filepath.Join(storageDir, filename)
}

func ToOutFilePath(filename string) string {
	storageDir := "./storage/out"
	return filepath.Join(storageDir, filename)
}

func GetFileSize(filename string) (string, error) {
	fileInfo, err := os.Stat(filename)
	if err != nil {
		return 0, err
	}
	FileSize, err := strconv.Atoi(fileInfo.Size())
	return FileSize, nil
}
