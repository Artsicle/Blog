package controllers

import (
	"blog/app/models"
	"blog/app/routes"
	"fmt"
	"github.com/revel/revel"
)

type App struct {
	Common
}

func (c App) Index() revel.Result {
	posts := &[]models.Post{}
	if err := models.FindAll(posts, false); err != nil {
		c.Flash.Error(err.Error())
		return c.Render()
	}
	return c.Render(posts)
}

func (c App) Register() revel.Result {
	return c.Render()
}

func (c App) SignIn() revel.Result {
	return c.Render()
}

func (c App) FindUser(email string) revel.Result {
	var user models.User
	attrs := map[string]interface{}{"email": email}

	if err := models.FindByMap(attrs, &user, true); err != nil {
		c.Flash.Error(err.Error())
		return c.Redirect(routes.App.SignIn())
	}
	c.Session["user"] = email
	return c.Redirect(routes.Posts.Index())
}

func (c App) CreateUser(user models.User) revel.Result {

	if err := models.CreateRecord(&user); err != nil {
		c.Flash.Error(err.Error())
		return c.Redirect(routes.App.Register())
	}
	c.Flash.Success(fmt.Sprintf("Welcome, %v %v",
		user.FirstName, user.LastName))
	c.Session["user"] = user.Email
	return c.Redirect(routes.Posts.Index())
}
