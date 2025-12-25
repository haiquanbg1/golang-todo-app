package models

import (
	"time"

	"gorm.io/gorm"
)

type Todo struct {
	ID          uint           `json:"id" gorm:"primaryKey;autoIncrement"`
	UserID      uint           `json:"user_id" gorm:"not null;index:idx_todos_user_id"`
	Task        string         `json:"task" gorm:"size:255;not null"`
	Description string         `json:"description" gorm:"type:text"`
	Status      TodoStatus     `json:"status" gorm:"type:varchar(20);not null;default:'pending';check:check_todos_status,status IN ('pending','in_progress','done')"`
	CreatedAt   time.Time      `json:"created_at" gorm:"not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt   time.Time      `json:"updated_at" gorm:"not null;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index:idx_todos_deleted_at"`

	User User `json:"user,omitempty" gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
}
