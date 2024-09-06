package cockroach_repositories

import (
	cockroach_database "github.com/axel-andrade/opina-ai-api/internal/adapters/secondary/database/cockroach"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type BaseCockroachRepository struct {
	Db *gorm.DB
	Tx *gorm.DB
}

func BuildBaseRepository() *BaseCockroachRepository {
	db := cockroach_database.GetDB()
	return &BaseCockroachRepository{Db: db, Tx: nil}
}

func (r *BaseCockroachRepository) getQueryOrTx() *gorm.DB {
	if r.Tx != nil {
		return r.Tx
	}

	return r.Db
}

func (r *BaseCockroachRepository) StartTransaction() error {
	// Note the use of tx as the database handle once you are within a transaction
	tx := r.Db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	r.Tx = tx

	return nil
}

func (r *BaseCockroachRepository) CommitTransaction() error {
	err := r.Tx.Commit().Error
	r.Tx = nil
	if err != nil {
		return err
	}
	return nil
}

func (r *BaseCockroachRepository) CancelTransaction() error {
	err := r.Tx.Rollback().Error
	r.Tx = nil

	if err != nil {
		return err
	}

	return nil
}

func (r *BaseCockroachRepository) NextEntityID() string {
	return uuid.NewV4().String()
}
