package cockroach_models

type ImportModel struct {
	BaseModel
	UserID       string  `gorm:"type:uuid;not null" json:"user_id"`
	Filename     string  `gorm:"type:varchar(255);not null" json:"filename"`
	Description  string  `gorm:"type:varchar(255);not null" json:"description"`
	Status       string  `gorm:"type:varchar(20);not null" json:"status"`
	TotalRecords int     `gorm:"type:integer;not null" json:"total_records"`
	ErrorMessage *string `gorm:"type:text" json:"error_message"`
}

func (ImportModel) TableName() string {
	return "imports"
}
