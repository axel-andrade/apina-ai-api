package cockroach_migrations

import (
	"log"

	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	if err := MigrateCreateVoterTable(db); err != nil {
		RollbackCreateVoterTable(db)
		log.Fatal(err)
	}
}
