package models

import (
	"time"
)

type Comment struct {
	Id        int64
	Body      string
	Commenter string
	PostId    int // Foreign key for Post (belongs to)
	CreatedAt time.Time
	UpdatedAt time.Time
}
