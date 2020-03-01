package models

import "time"

const (
	UnDone     = 0
	Done       = 1
	InProgress = 2
)

type TodoItem struct {
	ID        uint64    `gorm:"primary_key;auto_increment" json:"id"`
	User      User      `json:"user"`
	UserId    uint32    `gorm:"not null" json:"user_id"`
	Content   string    `gorm:"text;not null" json:"title"`
	Status    int       `gorm:"not null" json:"status"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}
