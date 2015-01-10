package controllers

import (
	"goblog/app/models"
)

type App struct {
	GormController
	CurrentUser *models.User
}

