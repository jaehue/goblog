package controllers

import (
	"code.google.com/p/go.crypto/bcrypt"
	"github.com/revel/revel"
	"goblog/app/models"
)

type User struct {
	App
}

func (c User) CheckUser() revel.Result {
	switch c.MethodName {
	case "Login", "CreateSession":
		return nil
	}

	if c.CurrentUser == nil {
		c.Flash.Error("Please log in first")
		return c.Redirect(User.Login)
	}

	return nil
}

func (c User) Edit() revel.Result {
	user := c.CurrentUser
	return c.Render(user)
}

func (c User) Update(name, oldPassword, newPassword, newPasswordConfirm string) revel.Result {
	if err := bcrypt.CompareHashAndPassword(c.CurrentUser.Password, []byte(oldPassword)); err != nil {
		c.Flash.Error("Old password isn't valid.")
		return c.Redirect(User.Edit)
	}

	var user models.User
	c.Txn.First(&user, c.CurrentUser.Id)
	user.Name = name

	if newPassword != "" && newPasswordConfirm != "" {
		if newPassword == newPasswordConfirm {
			bcryptPassword, _ := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
			user.Password = bcryptPassword
		} else {
			c.Flash.Error("Password doesn't match the confirmation.")
			return c.Redirect(User.Edit)
		}
	}

	c.Txn.Save(&user)
	return c.Redirect(Home.Index)
}

func (c User) Login() revel.Result {
	return c.Render()
}

func (c User) CreateSession(username, password string) revel.Result {
	var user models.User
	c.Txn.Where(&models.User{Username: username}).First(&user)

	err := bcrypt.CompareHashAndPassword(user.Password, []byte(password))
	if err == nil {
		authKey := revel.Sign(user.Username)
		c.Session["authKey"] = authKey
		c.Session["username"] = user.Username
		c.Session["userId"] = string(user.Id)
		if user.Role == "admin" {
			c.Session["isAdmin"] = "true"
		}
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

func (c User) DestroySession() revel.Result {
	// clear session
	for k := range c.Session {
		delete(c.Session, k)
	}
	return c.Redirect(Home.Index)
}
