package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID        uuid.UUID `gorm:"primaryKey;type:uuid" json:"id"`
	Name      string    `gorm:"size:20;not null;"    json:"name"`
	Role      string    `gorm:"size:10;default user" json:"role"`
	Key       Key
	CreatedAt time.Time `gorm:"autoCreateTime"       json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime:milli" json:"updated_at"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New()
	return
}

type Key struct {
	ID        uuid.UUID `gorm:"primaryKey;type:uuid" json:"id"`
	UserID    uuid.UUID `                            json:"user_id"`
	CreatedAt time.Time `gorm:"autoCreateTime"       json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime:milli" json:"updated_at"`
}

func (k *Key) BeforeCreate(tx *gorm.DB) (err error) {
	k.ID = uuid.New()
	return
}
