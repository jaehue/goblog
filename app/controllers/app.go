package controllers

import (
	"code.google.com/p/go.crypto/bcrypt"
	"github.com/revel/revel"
	"goblog/app/models"
)

type App struct {
	GormController
}

func (c App) Login() revel.Result {
	return c.Render()
}

func (c App) CreateSession(username, password string) revel.Result {
	var user models.User
	c.Txn.Where(&models.User{Username: username}).First(&user)

	err := bcrypt.CompareHashAndPassword(user.Password, []byte(password))
	if err == nil {
		c.Session["username"] = user.Username
		c.Session["role"] = user.Role
		c.Session["name"] = user.Name
		c.Flash.Success("Welcome, " + user.Name)
		return c.Redirect(Post.Index)
	}

	// clear session
	for k := range c.Session {
		delete(c.Session, k)
	}
	c.Flash.Out["username"] = username
	c.Flash.Error("Login failed")
	return c.Redirect(Home.Index)
}
