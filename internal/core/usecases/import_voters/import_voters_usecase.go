package import_voters

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"log"
	"sync"

	"github.com/axel-andrade/opina-ai-api/internal/core/domain"
	err_msg "github.com/axel-andrade/opina-ai-api/internal/core/domain/constants/errors"
)

type ImportVotersUC struct {
	Gateway ImportVotersGateway
}

func BuildImportVotersUC(g ImportVotersGateway) *ImportVotersUC {
	return &ImportVotersUC{g}
}

func (bs *ImportVotersUC) Execute(input ImportVotersInput) error {
	log.Println("Importing voters")

	// Use a wait group to wait for the goroutine to finish
	var wg sync.WaitGroup
	var err error
	// Buffered channel to hold the error
	errChan := make(chan error, 1)

	// Start a goroutine to execute the import
	wg.Add(1)
	go func() {
		// Defer the wait group to be marked as done when the goroutine finishes
		defer wg.Done()

		log.Println("Parsing CSV to domain")
		voters, parseErr := bs.parseCSVToDomain(input.Data)
		if parseErr != nil {
			errChan <- parseErr
			return
		}

		log.Println("Getting voters by cellphones")
		var votersCellphones []string

		for _, voter := range voters {
			votersCellphones = append(votersCellphones, voter.Cellphone)
		}

		log.Println("Checking existing voters")
		existingVoters, _ := bs.Gateway.GetVotersByCellphones(votersCellphones)
		existingVotersMap := make(map[string]*domain.Voter)

		for _, voter := range existingVoters {
			existingVotersMap[voter.Cellphone] = voter
		}

		var votersToCreate []*domain.Voter

		for _, voter := range voters {
			if _, exists := existingVotersMap[voter.Cellphone]; !exists {
				votersToCreate = append(votersToCreate, voter)
			}
		}

		if len(votersToCreate) == 0 {
			log.Println("No voters to create")
			return
		}

		log.Println("Creating voters")
		if createErr := bs.Gateway.CreateVoters(votersToCreate); createErr != nil {
			errChan <- createErr
		}
	}()

	wg.Wait()
	close(errChan)

	// Check if there was an error during the goroutine execution
	if err = <-errChan; err != nil {
		return err
	}

	return nil
}

func (bs *ImportVotersUC) parseCSVToDomain(data []byte) ([]*domain.Voter, error) {
	// Create a CSV reader
	reader := csv.NewReader(bytes.NewReader(data))

	// Read reads one record (a slice of fields) from r. The record is a slice of strings with each string representing one field.
	header, err := reader.Read()
	if err != nil {
		return nil, err
	}

	// Check if required fields are present
	requiredFields := map[string]bool{
		"full_name": false,
		"cellphone": false,
	}

	for _, field := range header {
		if _, exists := requiredFields[field]; exists {
			requiredFields[field] = true
		}
	}

	for _, found := range requiredFields {
		if !found {
			return nil, fmt.Errorf(err_msg.MISSING_REQUIRED_FIELDS_CSV)
		}
	}

	var voters []*domain.Voter

	// Iterates over the remaining lines
	for {
		record, err := reader.Read()
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			return nil, err
		}

		// Map the fields indexes
		dataMap := make(map[string]string)
		for i, field := range header {
			dataMap[field] = record[i]
		}

		// Create a new voter using the CSV data
		fullName := dataMap["full_name"]
		cellphone := dataMap["cellphone"]

		voter, err := domain.BuildNewVoter(fullName, cellphone)
		if err != nil {
			return nil, err
		}

		voters = append(voters, voter)
	}

	return voters, nil
}
