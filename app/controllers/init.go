package controllers

import (
	"github.com/revel/revel"
	"goblog/app/models"
	"log"
)

func (c *App) setCurrentUser() revel.Result {
	defer func() {
		log.Println("start ", c.Action)

		if c.CurrentUser != nil {
			c.RenderArgs["currentUser"] = c.CurrentUser
			log.Printf("current user: %q", c.CurrentUser)
		} else {
			delete(c.RenderArgs, "currentUser")
		}
	}()

	username, ok := c.Session["username"]
	if !ok || username == "" {
		return nil
	}

	authKey, ok := c.Session["authKey"]
	if !ok || authKey == "" {
		return nil
	}

	if match := revel.Verify(username, authKey); match {
		var user models.User
		c.Txn.Where(&models.User{Username: username}).First(&user)
		if &user != nil {
			c.CurrentUser = &user
		}
	}
	return nil
}

func init() {
	revel.InterceptMethod((*App).setCurrentUser, revel.BEFORE)
	revel.InterceptMethod(User.CheckUser, revel.BEFORE)
	revel.InterceptMethod(Post.CheckUser, revel.BEFORE)
	revel.InterceptMethod(Comment.CheckUser, revel.BEFORE)
}
