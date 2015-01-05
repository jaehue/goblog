package controllers

import (
	"github.com/revel/revel"
	"github.com/revel/revel/modules/db/app"
)

func InitDB() {
	db.Init()
	schema := `
CREATE TABLE IF NOT EXISTS posts (
    id          INTEGER PRIMARY KEY AUTOINCREMENT,
    title       TEXT                NOT NULL,
    body        TEXT                NOT NULL,
    created_at  DATETIME            NOT NULL,
    updated_at  DATETIME            NOT NULL
);
`
	db.Db.Exec(schema)
}

func init() {
	revel.OnAppStart(InitDB)
}
