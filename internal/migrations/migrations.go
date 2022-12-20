package migrations

import (
	"log"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"

	"go-test/internal/models"
	"go-test/pkg/database"
)

func Init() *gormigrate.Gormigrate {
	db := database.GetDB()
	m := gormigrate.New(db, gormigrate.DefaultOptions, []*gormigrate.Migration{
		Migrate20221128001(),
		// Migrate20221209001(),
	})

	m.InitSchema(func(tx *gorm.DB) error {
		err := tx.AutoMigrate(
			&models.User{},
			&models.Key{},
			// all other tables of you app
		)
		if err != nil {
			return err
		}

		// all other foreign keys...
		return nil
	})

	if err := m.Migrate(); err != nil {
		log.Fatalf("Could not migrate: %v", err)
	}
	// m.RollbackLast()
	return m
}
