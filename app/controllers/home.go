package controllers

import "github.com/revel/revel"

type Home struct {
	App
}

func (c Home) Index() revel.Result {
	return c.Render()
}
