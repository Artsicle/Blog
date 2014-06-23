package controllers

import (
	"Blog/app/helpers"
	"Blog/app/models"
	"github.com/revel/revel"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	var ulink, plink helpers.Linker
	ulink = &models.User{}
	plink = &models.Post{}
	h := &helpers.Helper{}
	return c.Render(h, ulink, plink)
}
