package controllers

import (
	"github.com/revel/revel"
	"goblog/app/models"
)

type App struct {
	*revel.Controller
}

func (c App) Login() revel.Result {
	return c.Render()
}

func (c App) Signin(username, password string) revel.Result {
	user := c.getUser(username)
	if user != nil && user.Password == password {
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

func (c App) getUser(username string) *models.User {
	switch username {
	case "admin":
		return &models.User{Name: "Admin", Username: "admin", Role: "admin", Password: "admin"}
	case "user":
		return &models.User{Name: "User", Username: "user", Role: "user", Password: "user"}
	default:
		return nil
	}
}
