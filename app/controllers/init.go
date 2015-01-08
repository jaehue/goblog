package controllers

import (
	"github.com/revel/revel"
)

func init() {
	revel.InterceptMethod(Post.CheckUser, revel.BEFORE)
	revel.InterceptMethod(Comment.CheckUser, revel.BEFORE)
}
