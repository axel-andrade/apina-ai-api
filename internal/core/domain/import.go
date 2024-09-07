package domain

const (
	ImportStatusProcessing = "processing"
	ImportStatusDone       = "done"
	ImportStatusError      = "error"
)

type Import struct {
	Base
	UserID       string  `json:"user_id"`
	Filename     string  `json:"filename"`
	Description  string  `json:"description"`
	Status       string  `json:"status"`
	TotalRecords int     `json:"total_records"`
	ErrorMessage *string `json:"error_message"`
}

func BuildNewImport(userId, filename, description string) *Import {
	return &Import{
		UserID:      userId,
		Filename:    filename,
		Description: description,
		Status:      ImportStatusProcessing,
	}
}
