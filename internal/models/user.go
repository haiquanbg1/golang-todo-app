package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint           `json:"id" gorm:"primaryKey;autoIncrement"`
	Username  string         `json:"username" gorm:"uniqueIndex:ui_users_username;size:255;not null"`
	Password  string         `json:"-" gorm:"size:255;not null"`
	CreatedAt time.Time      `json:"created_at" gorm:"not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"not null;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index:idx_users_deleted_at"`

	Todos []Todo `json:"todos" gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
}
