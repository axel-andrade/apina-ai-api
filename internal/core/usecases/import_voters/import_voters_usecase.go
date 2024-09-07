package import_voters

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"log"

	"github.com/axel-andrade/opina-ai-api/internal/core/domain"
	err_msg "github.com/axel-andrade/opina-ai-api/internal/core/domain/constants/errors"
)

type ImportVotersUC struct {
	Gateway ImportVotersGateway
}

func BuildImportVotersUC(g ImportVotersGateway) *ImportVotersUC {
	return &ImportVotersUC{g}
}

func (bs *ImportVotersUC) Execute(input ImportVotersInput) (*ImportVotersOutput, error) {
	log.Println("Creating import")
	createdImport, err := bs.Gateway.CreateImport(&domain.Import{
		UserID:   input.UserID,
		Filename: "voters.csv",
	})
	if err != nil {
		return nil, err
	}

	// Retorne imediatamente com o ID da importação
	go bs.processImport(createdImport, input.Data)

	return &ImportVotersOutput{Import: createdImport}, nil
}

func (bs *ImportVotersUC) processImport(createdImport *domain.Import, data []byte) {
	log.Println("Importing voters")
	var votersToCreate []*domain.Voter

	// Processamento existente movido para esta função
	voters, err := bs.parseCSVToDomain(data)
	if err != nil {
		log.Println("Error parsing CSV:", err)
		bs.updateImportError(createdImport, err)
		return
	}

	// Processamento de checagem de eleitores existentes
	var votersCellphones []string
	for _, voter := range voters {
		votersCellphones = append(votersCellphones, voter.Cellphone)
	}

	existingVoters, _ := bs.Gateway.GetVotersByCellphones(votersCellphones)
	existingVotersMap := make(map[string]*domain.Voter)

	for _, voter := range existingVoters {
		existingVotersMap[voter.Cellphone] = voter
	}

	for _, voter := range voters {
		if _, exists := existingVotersMap[voter.Cellphone]; !exists {
			votersToCreate = append(votersToCreate, voter)
		}
	}

	if len(votersToCreate) > 0 {
		if err := bs.Gateway.CreateVoters(votersToCreate); err != nil {
			bs.updateImportError(createdImport, err)
			return
		}
	}

	createdImport.TotalRecords = len(votersToCreate)
	createdImport.Status = domain.ImportStatusCompleted

	log.Println("Updating import")
	bs.Gateway.UpdateImport(createdImport)
}

func (bs *ImportVotersUC) updateImportError(createdImport *domain.Import, err error) {
	createdImport.Status = domain.ImportStatusError
	createdImport.ErrorMessage = err.Error()
	bs.Gateway.UpdateImport(createdImport)
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
