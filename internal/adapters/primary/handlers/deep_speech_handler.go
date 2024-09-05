package handlers

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/google/uuid"
)

const (
	modelPath      = "/opt/deepspeech/deepspeech-0.9.3-models.pbmm"
	scorerPath     = "/opt/deepspeech/deepspeech-0.9.3-models.scorer"
	deepspeechPath = "/usr/local/bin/deepspeech"
)

type DeepSpeechHandler struct{}

func BuildDeepSpeechHandler() *DeepSpeechHandler {
	return &DeepSpeechHandler{}
}

func (e *DeepSpeechHandler) TransformAudioToText(file []byte, filename string) (string, error) {
	// Create a temporary file
	tempFile, err := os.CreateTemp("", uuid.New().String()+filepath.Ext(filename))
	if err != nil {
		return "", fmt.Errorf("failed to create temp file: %w", err)
	}
	defer os.Remove(tempFile.Name()) // Clean up the temp file

	// Write audio bytes to the temporary file
	if _, err := tempFile.Write(file); err != nil {
		return "", fmt.Errorf("failed to write to temp file: %w", err)
	}
	tempFile.Close()

	// Determine if conversion to WAV is necessary
	var filePath string
	ext := filepath.Ext(filename)
	if ext != ".wav" {
		filePath = "/tmp/" + uuid.New().String() + ".wav"
		if err := convertToWAV(tempFile.Name(), filePath); err != nil {
			return "", fmt.Errorf("error converting %s to WAV: %w", ext, err)
		}
		defer os.Remove(filePath) // Clean up the converted file
	} else {
		filePath = tempFile.Name()
	}

	// Process the file with DeepSpeech
	transcription, err := processFile(filePath)
	if err != nil {
		return "", fmt.Errorf("failed to process file with DeepSpeech: %w", err)
	}

	return transcription, nil
}

func convertToWAV(inputPath, wavPath string) error {
	cmd := exec.Command("ffmpeg", "-i", inputPath, "-ar", "16000", "-ac", "1", wavPath)
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("ffmpeg conversion error: %w, stderr: %s", err, stderr.String())
	}
	return nil
}

func processFile(filePath string) (string, error) {
	fmt.Printf("Processing file %s\n", filePath)

	cmd := exec.Command(deepspeechPath,
		"--model", modelPath,
		"--scorer", scorerPath,
		"--audio", filePath)

	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		println(stderr.String())
		return "", fmt.Errorf("error processing file %s: %w, output: %s, error output: %s", filePath, err, out.String(), stderr.String())
	}

	return out.String(), nil
}
