package cockroach_migrations

import (
	"fmt"

	cockroach_models "github.com/axel-andrade/opina-ai-api/internal/adapters/secondary/database/cockroach/models"
	"gorm.io/gorm"
)

func MigrateCreateVoterTable(tx *gorm.DB) error {
	if !tx.Migrator().HasTable(&cockroach_models.VoterModel{}) {
		if err := tx.AutoMigrate(&cockroach_models.VoterModel{}); err != nil {
			return err
		}

		fmt.Println("Migration executed: 1683385982286_create_expenses")
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
