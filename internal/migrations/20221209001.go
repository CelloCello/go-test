package migrations

import (
	"go-test/internal/models"
	"log"
	"time"

	"github.com/go-gormigrate/gormigrate/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func Migrate20221209001() *gormigrate.Migration {
	log.Println("do migrate 20221209001")
	return &gormigrate.Migration{
		// add age column to users
		ID: "20221209001",
		Migrate: func(tx *gorm.DB) error {
			// when table already exists, it just adds fields as columns
			type Key struct {
				ID        uuid.UUID
				UserID    uuid.UUID
				CreatedAt time.Time
				UpdatedAt time.Time
			}
			type User struct {
				Key Key
			}
			if err := tx.AutoMigrate(&Key{}, &User{}); err != nil {
				return err
			}
			return nil
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.Migrator().DropTable(&models.Key{})
		},
	}
}
