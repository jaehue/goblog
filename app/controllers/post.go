package controllers

import (
	"github.com/revel/revel"
	"github.com/russross/blackfriday"
	"goblog/app/models"
	"goblog/app/routes"
	"html/template"
)

type Post struct {
	App
}

func (c Post) CheckUser() revel.Result {
	switch c.MethodName {
	case "Index", "Show":
		return nil
	}

	if c.CurrentUser == nil {
		c.Flash.Error("Please log in first")
		return c.Redirect(User.Login)
	}

	if c.CurrentUser.Role != "admin" {
		c.Response.Status = 401 // Unauthorized
		c.Flash.Error("You are not admin")
		return c.Redirect(User.Login)
	}
	return nil
}

func (c Post) Index() revel.Result {
	var posts []models.Post
	c.Txn.Order("created_at desc").Find(&posts)
	for i, p := range posts {
		posts[i].HtmlBody = template.HTML(string(blackfriday.MarkdownCommon([]byte(p.Body))))
	}
	return c.Render(posts)
}

func (c Post) Show(id int) revel.Result {
	var post models.Post
	c.Txn.First(&post, id)
	c.Txn.Where(&models.Comment{PostId: id}).Find(&post.Comments)
	post.HtmlBody = template.HTML(string(blackfriday.MarkdownCommon([]byte(post.Body))))
	return c.Render(post)
}

func (c Post) Update(id int, title, body string) revel.Result {
	var post models.Post
	c.Txn.First(&post, id)
	post.Title = title
	post.Body = body

	c.Txn.Save(&post)
	return c.Redirect(routes.Post.Show(id))
}

func (c Post) Create(title, body string) revel.Result {
	post := models.Post{Title: title, Body: body}
	c.Txn.Create(&post)
	return c.Redirect(routes.Post.Show(int(post.Id)))
}

func (c Post) New() revel.Result {
	post := models.Post{}
	return c.Render(post)
}

func (c Post) Edit(id int) revel.Result {
	var post models.Post
	c.Txn.First(&post, id)
	return c.Render(post)
}

func (c Post) Destroy(id int) revel.Result {
	c.Txn.Where("post_id = ?", id).Delete(&models.Comment{})
	c.Txn.Where("id = ?", id).Delete(&models.Post{})
	return c.Redirect(routes.Post.Index())
}
