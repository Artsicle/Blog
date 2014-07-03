package controllers

import (
	"blog/app/models"
	"fmt"
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
		ok, err := models.FindByMap(attrs, user)
		if !ok {
			fmt.Printf("Error: %s", err.Error())
		} else {
			return user
		}
	}
	return nil
}
