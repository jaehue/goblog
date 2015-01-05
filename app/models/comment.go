package models

import (
	"time"
)

type Comment struct {
	Id        int
	Body      string
	Commenter string
	PostId    int
	CreatedAt time.Time
	UpdatedAt time.Time
}
