package handlers

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
)

const (
	tempFilePattern = "audio-*.wav"
)

type FileHandler struct{}

func BuildFileHandler() *FileHandler {
	return &FileHandler{}
}

func (e *FileHandler) ConvertFileToBytes(file multipart.File) ([]byte, error) {
	tempFile, err := os.CreateTemp("/tmp", tempFilePattern)
	if err != nil {
		return nil, fmt.Errorf("unable to create temp file: %v", err)
	}

	defer tempFile.Close()

	if _, err = io.Copy(tempFile, file); err != nil {
		return nil, fmt.Errorf("unable to save file: %v", err)
	}

	file.Close()

	// Read the content of tempFile into fileBytes
	tempFile.Seek(0, 0)
	fileBytes, err := io.ReadAll(tempFile)

	if err != nil {
		return nil, fmt.Errorf("unable to read file: %v", err)
	}

	return fileBytes, nil
}
