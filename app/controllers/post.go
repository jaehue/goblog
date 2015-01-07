package controllers

import (
	"database/sql"
	"fmt"
	"github.com/revel/revel"
	"github.com/revel/revel/modules/db/app"
	"goblog/app/models"
	"goblog/app/routes"
	"time"
)

type Post struct {
	*revel.Controller
	db.Transactional
}

func (c Post) CheckUser() revel.Result {
	switch c.MethodName {
	case "Index", "Show":
		return nil
	}
	_, ok := c.Session["username"]
	if !ok {
		c.Flash.Error("Please log in first")
		return c.Redirect(App.Login)
	}
	return nil
}

func getPost(txn *sql.Tx, id int) (models.Post, error) {
	post := models.Post{}
	err := txn.QueryRow("select id, title, body, created_at, updated_at from posts where id=?", id).
		Scan(&post.Id, &post.Title, &post.Body, &post.CreatedAt, &post.UpdatedAt)

	switch {
	case err == sql.ErrNoRows:
		return post, fmt.Errorf("No post with that ID - %d.", id)
	case err != nil:
		return post, err
	}

	post.Comments = getComments(txn, id)

	return post, nil
}

func getComments(txn *sql.Tx, postId int) (comments []models.Comment) {
	rows, err := txn.Query("select id, body, commenter, post_id, created_at, updated_at from comments where post_id=? order by created_at desc", postId)
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		comment := models.Comment{}
		if err := rows.Scan(&comment.Id, &comment.Body, &comment.Commenter, &comment.PostId, &comment.CreatedAt, &comment.UpdatedAt); err != nil {
			panic(err)
		}
		comments = append(comments, comment)
	}
	return
}

func (c Post) Index() revel.Result {
	var posts []models.Post
	rows, err := c.Txn.Query("select id, title, body, created_at, updated_at from posts order by created_at desc")
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		post := models.Post{}
		if err := rows.Scan(&post.Id, &post.Title, &post.Body, &post.CreatedAt, &post.UpdatedAt); err != nil {
			panic(err)
		}
		posts = append(posts, post)
	}

	return c.Render(posts)
}

func (c Post) Show(id int) revel.Result {
	post, err := getPost(c.Txn, id)
	if err != nil {
		panic(err)
	}

	return c.Render(post)
}

func (c Post) Update(id int, title, body string) revel.Result {
	if _, err := c.Txn.Exec("update posts set title=?, body=?, updated_at=? where id=?", title, body, time.Now(), id); err != nil {
		panic(err)
	}
	return c.Redirect(routes.Post.Show(id))
}

func (c Post) Create(title, body string) revel.Result {
	result, err := c.Txn.Exec("insert into posts(title, body, created_at, updated_at) values(?,?,?,?)", title, body, time.Now(), time.Now())
	id, err := result.LastInsertId()

	if err != nil {
		panic(err)
	}

	return c.Redirect(routes.Post.Show(int(id)))
}

func (c Post) New() revel.Result {
	post := models.Post{}
	return c.Render(post)
}

func (c Post) Edit(id int) revel.Result {
	post, err := getPost(c.Txn, id)
	if err != nil {
		panic(err)
	}

	return c.Render(post)
}

func (c Post) Destroy(id int) revel.Result {
	if _, err := c.Txn.Exec("delete from posts where id=?", id); err != nil {
		panic(err)
	}
	return c.Redirect(routes.Post.Index())
}
