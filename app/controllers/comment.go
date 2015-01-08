package controllers

import (
	"github.com/revel/revel"
	"goblog/app/models"
	"goblog/app/routes"
)

type Comment struct {
	GormController
}

func (c Comment) CheckUser() revel.Result {
	_, ok := c.Session["username"]
	if !ok {
		c.Flash.Error("Please log in first")
		return c.Redirect(App.Login)
	}
	return nil
}

func (c Comment) Create(postId int, body, commenter string) revel.Result {
	comment := models.Comment{PostId: postId, Body: body, Commenter: commenter}
	c.Txn.Create(&comment)
	return c.Redirect(routes.Post.Show(postId))
}

func (c Comment) Destroy(postId, id int) revel.Result {
	c.Txn.Where("id = ?", id).Delete(&models.Comment{})
	return c.Redirect(routes.Post.Show(postId))
}
