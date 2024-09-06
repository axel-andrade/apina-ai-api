package cockroach_migrations

import (
	"log"

	cockroach_models "github.com/axel-andrade/opina-ai-api/internal/adapters/secondary/database/cockroach/models"
	"gorm.io/gorm"
)

func MigrateCreateVoterTable(tx *gorm.DB) error {
	if !tx.Migrator().HasTable(&cockroach_models.VoterModel{}) {
		if err := tx.AutoMigrate(&cockroach_models.VoterModel{}); err != nil {
			return err
		}
		log.Printf("Migration executed: 20240904205730_create_voters_table")
	}

	return nil
}

func RollbackCreateVoterTable(tx *gorm.DB) error {
	err := tx.Migrator().DropTable(&cockroach_models.VoterModel{})
	if err != nil {
		return err
	}

	return nil
}
