package models

import "time"

// Article model
type Article struct {
	ID        int        `json:"id"`
	UserID    int        `json:"user_id" gorm:"index"`
	Title     string     `json:"title"`
	Content   string     `json:"content"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at" sql:"index"`
}
