package models

import "time"

type TodoList struct {
	ID          uint64    `gorm:"primary_key;auto_increment" json:"id"`
	User        User      `json:"user"`
	UserId      uint32    `gorm:"not null" json:"user_id"`
	Title       string    `gorm:"text;not null" json:"title"`
	Description string    `gorm:"text;not null" json:"description"`
	CreatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}
