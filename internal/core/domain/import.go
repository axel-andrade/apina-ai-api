package domain

const (
	ImportStatusProcessing = "processing"
	ImportStatusCompleted  = "completed"
	ImportStatusError      = "error"
)

type Import struct {
	Base
	UserID       string `json:"user_id"`
	Filename     string `json:"filename"`
	Status       string `json:"status"`
	TotalRecords int    `json:"total_records"`
	ErrorMessage string `json:"error_message"`
}

func BuildNewImport(userId, filename string) *Import {
	return &Import{
		UserID:   userId,
		Filename: filename,
		Status:   ImportStatusProcessing,
	}
}
