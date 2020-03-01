package models

import (
	"errors"
	"html"
	"strings"
	"time"
)

type TodoList struct {
	ID          uint64    `gorm:"primary_key;auto_increment" json:"id"`
	User        User      `json:"user"`
	UserId      uint32    `gorm:"not null" json:"user_id"`
	Title       string    `gorm:"text;not null" json:"title"`
	Description string    `gorm:"text;not null" json:"description"`
	CreatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (tl *TodoList) Prepare() {
	tl.Title = html.EscapeString(strings.TrimSpace(tl.Title))
	tl.Description = html.EscapeString(strings.TrimSpace(tl.Description))
	tl.User = User{}
	tl.CreatedAt = time.Now()
	tl.UpdatedAt = time.Now()
}

func (tl *TodoList) Validate() map[string]string {
	var err error
	var errorMessages = make(map[string]string)

	if tl.Title == "" {
		err = errors.New("required title")
		errorMessages["Required_title"] = err.Error()
	}
	if tl.Description == "" {
		err = errors.New("required description")
		errorMessages["Required_description"] = err.Error()
	}
	if tl.UserId < 1 {
		err = errors.New("required user")
		errorMessages["Required_user"] = err.Error()
	}

	return errorMessages
}

func (tl *TodoList) CreateTodoList() (*TodoList, error) {
	return tl, nil
}

func (tl *TodoList) FindAllTodoLists() (*[]TodoList, error) {
	// TODO use database!
	todoLists := []TodoList{
		{
			ID: 1,
			User: User{
				ID:         11,
				Username:   "Selim",
				Email:      "selim@gmail.com",
				Password:   "123",
				AvatarPath: "1212231",
			},
			UserId:      1,
			Title:       "Todo Item 1",
			Description: "Do Something",
		},
	}
	return &todoLists, nil
}
