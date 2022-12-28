package book

import (
	"database/sql"
	"time"
)

type Book struct {
	ID        uint         `json:"id" gorm:"primarykey"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt time.Time    `json:"updated_at"`
	DeletedAt sql.NullTime `json:"deleted_at" gorm:"index"`
	Title     string       `json:"title"`
	Author    string       `json:"author"`
	Rating    int          `json:"rating"`
}
