package controllers

import (
	"github.com/revel/revel"
	"github.com/revel/revel/modules/db/app"
	"goblog/app/routes"
	"time"
)

type Comment struct {
	*revel.Controller
	db.Transactional
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
	_, err := c.Txn.Exec("insert into comments(body, commenter, post_id, created_at, updated_at) values(?,?,?,?,?)", body, commenter, postId, time.Now(), time.Now())
	if err != nil {
		panic(err)
	}

	return c.Redirect(routes.Post.Show(postId))
}

func (c Comment) Destroy(postId, id int) revel.Result {
	if _, err := c.Txn.Exec("delete from comments where id=?", id); err != nil {
		panic(err)
	}
	return c.Redirect(routes.Post.Show(postId))
}
