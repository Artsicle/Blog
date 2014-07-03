package controllers

import (
	"blog/app/routes"
	"github.com/revel/revel"
)

type Posts struct {
	Common
}

func (p Posts) checkUser() revel.Result {
	if user := p.connected(); user == nil {
		p.Flash.Error("Please log in first")
		return p.Redirect(routes.App.Index())
	}
	return nil
}

func (p Posts) Index() revel.Result {
	return p.Render()
}
