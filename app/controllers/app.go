package controllers

import (
	. "blog/app/db"
	"blog/app/helpers"
	"blog/app/models"
	"blog/app/routes"
	"fmt"
	"github.com/revel/revel"
)

type App struct {
	*revel.Controller
}

func (c App) AddUser() revel.Result {
	if user := c.connected(); user != nil {
		c.RenderArgs["user"] = user
	}
	return nil
}

func (c App) connected() *models.User {
	if c.RenderArgs["user"] != nil {
		return c.RenderArgs["user"].(*models.User)
	}
	return nil
}

func (c App) Index() revel.Result {
	var ulink, plink helpers.Linker
	ulink = &models.User{}
	plink = &models.Post{}
	h := &helpers.Helper{}
	return c.Render(h, ulink, plink)
}

func (c App) Register() revel.Result {
	return c.Render()
}

func (c App) CreateUser(user models.User) revel.Result {
	if err := DB.Save(&user).Error; err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("User created")
	}

	c.Flash.Success(fmt.Sprintf("Welcome, %v %v",
		user.FirstName, user.LastName))

	return c.Redirect(routes.App.Index())
}
