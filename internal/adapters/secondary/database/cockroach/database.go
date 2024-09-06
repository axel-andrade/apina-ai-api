package cockroach_database

import (
	"log"
	"os"
	"time"

	cockroach_migrations "github.com/axel-andrade/opina-ai-api/internal/adapters/secondary/database/cockroach/migrations"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func ConnectDB() {
	dsn := os.Getenv("DSN_COCKROACH")
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	db = database

	config, _ := db.DB()

	config.SetMaxIdleConns(10)
	config.SetMaxOpenConns(100)
	config.SetConnMaxLifetime(time.Hour)

	if os.Getenv("DB_AUTO_MIGRATE") == "true" {
		cockroach_migrations.RunMigrations(db)
	}

	// Com o defer o go vai conseguir identificar quando executar uma determinada ação
	// defer config.Close()
}

func CloseDB() error {
	config, err := db.DB()
	if err != nil {
		return err
	}

	err = config.Close()
	if err != nil {
		return err
	}

	return nil
}

func GetDB() *gorm.DB {
	return db
}
