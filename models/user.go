package models

import (
	"time"
)

type User struct {
	ID        int        `json:"id" gorm:"primaryKey"`
	Name      string     `json:"name"`
	Email     string     `json:"email" gorm:"uniqueIndex"`
	Password  string     `json:"-"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at" sql:"index"`
	Biodata   *Biodata   `json:"biodata" gorm:"foreignKey:UserID"`
}

func (User) TableName() string {
	return "users"
}
