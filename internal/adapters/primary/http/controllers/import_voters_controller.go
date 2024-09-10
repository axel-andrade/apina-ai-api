package controllers

import (
	"io"

	"github.com/axel-andrade/opina-ai-api/internal/adapters/primary/http/presenters"
	"github.com/axel-andrade/opina-ai-api/internal/core/usecases/import_voters"
	"github.com/gin-gonic/gin"
)

type ImportVotersController struct {
	Usecase   import_voters.ImportVotersUC
	Presenter presenters.ImportVotersPresenter
}

func BuildImportVotersController(uc *import_voters.ImportVotersUC, ptr *presenters.ImportVotersPresenter) *ImportVotersController {
	return &ImportVotersController{Usecase: *uc, Presenter: *ptr}
}

// @Summary		Import voters
// @Description	Import voters from a file
// @Tags			voters
// @Accept			json
// @Produce		json
// @Success		204
// @Router			/api/v1/voters/import [post]
func (ctrl *ImportVotersController) Handle(c *gin.Context) {
	userID := "123"

	// Get the file from the form-data
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(400, gin.H{"error": "Failed to get file"})
		return
	}

	// Open the file
	openedFile, err := file.Open()
	if err != nil {
		c.JSON(400, gin.H{"error": "Failed to open file"})
		return
	}
	defer openedFile.Close()

	// Read the file content to transform into []byte
	fileBytes, err := io.ReadAll(openedFile)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to read file"})
		return
	}

	var input import_voters.ImportVotersInput
	input.UserID = userID
	input.Filename = file.Filename
	input.Data = fileBytes

	result, err := ctrl.Usecase.Execute(input)
	output := ctrl.Presenter.Show(result, err)

	c.JSON(output.StatusCode, output.Data)
}
