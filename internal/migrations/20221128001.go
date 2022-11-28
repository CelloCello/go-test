package migrations

import (
	"go-test/internal/models"
	"log"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func Migrate20221128001() *gormigrate.Migration {
	log.Println("do migrate 20221128001")
	return &gormigrate.Migration{
		// add age column to users
		ID: "20221128001",
		Migrate: func(tx *gorm.DB) error {
			// when table already exists, it just adds fields as columns
			type User struct {
				Role string
			}
			if err := tx.AutoMigrate(&User{}); err != nil {
				return err
			}
			tx.Model(&models.User{}).Where("role is NULL").Updates(models.User{Role: "user"})
			return nil
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.Migrator().DropColumn(&models.User{}, "role")
		},
	}
}
