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
	var ulink models.Linker
	ulink = &models.User{}
	return c.Render(ulink)
}

func (c App) Register() revel.Result {
	return c.Render()
}

func (c App) CreateUser(user models.User) revel.Result {
	ok, err := models.CreateRecord(&user)
	if !ok {
		c.Flash.Error(err.Error())
		return c.Redirect(routes.App.Register())
	} else {
		c.Flash.Success(fmt.Sprintf("Welcome, %v %v",
			user.FirstName, user.LastName))
		c.Session["user"] = user.Email
		return c.Redirect(routes.Posts.Index())
	}
}
