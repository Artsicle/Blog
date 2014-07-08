package controllers

import (
	"blog/app/models"
	"blog/app/routes"
	"github.com/revel/revel"
)

type Posts struct {
	Common
}

func (p Posts) checkUser() revel.Result {
	if user := p.connected(); user == nil {
		p.Flash.Error("Please log in first")
		return p.Redirect(routes.Pages.Index())
	}
	return nil
}

func (p Posts) Index() revel.Result {
	user := p.connected()
	query := map[string]interface{}{"user_id": user.Id}
	posts := &[]models.Post{}
	if err := models.FindByMap(query, posts, false); err != nil {
		p.Flash.Error(err.Error())
		return p.Redirect(routes.Pages.Index())
	}
	return p.Render(posts)
}

func (p Posts) New() revel.Result {
	return p.Render()
}

func (p Posts) CreatePost(post models.Post) revel.Result {
	if err := models.CreateRecord(&post); err != nil {
		p.Flash.Error(err.Error())
		return p.Redirect(routes.Posts.New())
	}
	p.Flash.Success("New Post created Successfully!")
	return p.Redirect(routes.Posts.Index())
}

func (p Posts) Show(id int) revel.Result {
	query := map[string]interface{}{"id": id}
	post := &models.Post{}
	if err := models.FindByMap(query, post, true); err != nil {
		p.Flash.Error(err.Error())
		return p.Redirect(routes.Posts.Index())
	}
	return p.Render(post)
}
