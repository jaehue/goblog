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
	return post, nil
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