package models

import (
	"time"
)

type Post struct {
	Id        int64
	Title     string
	Body      string
	CreatedAt time.Time
	UpdatedAt time.Time
	Comments  []Comment // One-To-Many relationship (has many)
}
