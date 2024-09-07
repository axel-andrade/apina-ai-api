package common_ptr

import (
	"github.com/axel-andrade/opina-ai-api/internal/core/domain"
)

type ImportFormatted struct {
	ID           string `json:"id" example:"123" description:"O ID do contéudo"`
	UserID       string `json:"user_id" example:"123" description:"O ID do usuário que importou o arquivo"`
	Filename     string `json:"file_name" example:"voters.csv" description:"O nome do arquivo importado"`
	Description  string `json:"description" example:"Voters list" description:"A descrição do arquivo importado"`
	Status       string `json:"status" example:"completed" description:"O status do arquivo importado"`
	TotalRecords int    `json:"total_records" example:"100" description:"O total de registros importados"`
	ErrorMessage string `json:"error_message" example:"Error message" description:"A mensagem de erro do arquivo importado"`
	CreatedAt    string `json:"created_at" example:"2021-01-01T00:00:00Z" description:"Data de criação do contéudo"`
	UpdatedAt    string `json:"updated_at" example:"2021-01-01T00:00:00Z" description:"Data de atualização do contéudo"`
}

type ImportPresenter struct{}

func BuildImportPresenter() *ImportPresenter {
	return &ImportPresenter{}
}

func (ptr *ImportPresenter) Format(i *domain.Import) ImportFormatted {
	return ImportFormatted{
		ID:           i.ID,
		UserID:       i.UserID,
		Filename:     i.Filename,
		Status:       i.Status,
		TotalRecords: i.TotalRecords,
		ErrorMessage: i.ErrorMessage,
		CreatedAt:    i.CreatedAt.Format("2006-01-02T15:04:05Z"),
		UpdatedAt:    i.UpdatedAt.Format("2006-01-02T15:04:05Z"),
	}
}

func (ptr *ImportPresenter) FormatList(imports []domain.Import) []ImportFormatted {
	var importsFormatted []ImportFormatted = make([]ImportFormatted, 0)

	for _, i := range imports {
		importsFormatted = append(importsFormatted, ptr.Format(&i))
	}

	return importsFormatted
}
