package models

import (
	"time"
)

type Post struct {
	Id        int
	Title     string
	Body      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
