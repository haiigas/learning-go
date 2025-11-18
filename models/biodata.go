package models

import (
	"time"
)

type Biodata struct {
	ID        int        `json:"id" gorm:"primaryKey"`
	UserID    int        `json:"user_id" gorm:"uniqueIndex"`
	Phone     string     `json:"phone"`
	Address   string     `json:"address"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at" sql:"index"`
}

func (Biodata) TableName() string {
	return "biodatas"
}
