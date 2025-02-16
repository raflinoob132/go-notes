package models

import "time"

//CRUD MODEL
type PostPayload struct {
	Title    string `json:"title"`
	Content  string `json:"content"`
	Category string `json:"category,omitempty"`
}

type Note struct {
	ID        uint      `gorm:"primaryKey;autoIncrement;not null" json:"id,omitempty"`
	Title     string    `gorm:"type:varchar(255);unique;not null" json:"title"`
	Content   string    `json:"content"`
	Category  string    `json:"category,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type SearchPayload struct {
	Title       string `json:"title"`
	SearchAll   string `json:"search_all"`
	Page        int    `json:"page"`
	Limit       int    `json:"limit"`
	SearchTitle string `json:"search_title"`
	SearchCat   string `json:"search_cat"`
}

//MODEL USER
type User struct {
	ID        int       `json:"id" gorm:"primaryKey"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
}

type LoginPayload struct {
	UsernameOrEmail string `json:"username_or_email"`
	Password        string `json:"password"`
}

type Favorite struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	UserID    uint      `gorm:"not null" json:"user_id"`
	NoteID   uint      `gorm:"not null" json:"note_id`
	CreatedAt time.Time `gorm:"autoCreateTime"`
}
