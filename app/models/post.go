package models

import (
	"html/template"
	"time"
)

type Post struct {
	Id        int64
	Title     string
	Body      string
	CreatedAt time.Time
	UpdatedAt time.Time
	Comments  []Comment     // One-To-Many relationship (has many)
	HtmlBody  template.HTML `sql:"-"`
}
