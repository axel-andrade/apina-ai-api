package process_audio

import (
	"fmt"
	"log"

	"github.com/axel-andrade/opina-ai-api/internal/core/domain"
	err_msg "github.com/axel-andrade/opina-ai-api/internal/core/domain/constants/errors"
)

type ProcessAudioUC struct {
	Gateway ProcessAudioGateway
}

func BuildProcessAudioUC(g ProcessAudioGateway) *ProcessAudioUC {
	return &ProcessAudioUC{g}
}

func (bs *ProcessAudioUC) Execute(input ProcessAudioInput) error {
	// Fetch content by ID
	content, err := bs.Gateway.FindContentByID(input.ContentID)

	if err != nil {
		return fmt.Errorf(err_msg.INTERNAL_SERVER_ERROR)
	}

	if content == nil {
		return fmt.Errorf(err_msg.CONTENT_NOT_FOUND)
	}

	if content.Status != domain.ContentStatusProcessing {
		return fmt.Errorf(err_msg.CONTENT_ALREADY_PROCESSED)
	}

	// Respond with success immediately and process in background
	go func() {
		// Background processing
		log.Println("Processing content in background:", content.ID)
		text, err := bs.Gateway.TransformAudioToText(content.File, content.Filename)
		if err != nil {
			// Handle error, log it, etc.
			fmt.Println("Error transforming audio to text:", err)
			return
		}

		if ct, err := bs.Gateway.CorrectText(text); err == nil {
			text = ct
		} else {
			fmt.Println("Error correcting text:", err)
		}

		// Update content with transformed text and status
		log.Println("Updating content with transformed text:", content.ID)

		contentToUpdate := domain.Content{
			Text:   text,
			Status: domain.ContentStatusReady,
		}

		if err = bs.Gateway.UpdateContent(content.ID, &contentToUpdate); err != nil {
			// Handle error, log it, etc.
			fmt.Println("Error updating content with transformed text:", err)
			return
		}

		// Success in background processing
		fmt.Println("Content processed successfully:", content.ID)
	}()

	return nil
}
