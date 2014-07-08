package controllers

import (
	"blog/app/models"
	"github.com/revel/revel"
)

type Pages struct {
	Common
}

func (p Pages) Index() revel.Result {
	posts := &[]models.Post{}
	if err := models.FindAll(posts, false); err != nil {
		p.Flash.Error(err.Error())
		return p.Render()
	}
	return p.Render(posts)
}

func (p Pages) About() revel.Result {
	return p.Render()
}

func (p Pages) Team() revel.Result {
	return p.Render()
}
