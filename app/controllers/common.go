package controllers

import (
	"blog/app/models"
	"github.com/revel/revel"
)

type Common struct {
	*revel.Controller
}

func (c Common) AddUser() revel.Result {
	if user := c.connected(); user != nil {
		c.RenderArgs["user"] = user
	}
	return nil
}

func (c Common) connected() *models.User {
	if c.RenderArgs["user"] != nil {
		return c.RenderArgs["user"].(*models.User)
	}
	if email, ok := c.Session["user"]; ok {
		attrs := map[string]interface{}{"email": email}
		user := &models.User{}

		if err := models.FindByMap(attrs, user, true); err != nil {
			return nil
		}
		return user

	}
	return nil
}
