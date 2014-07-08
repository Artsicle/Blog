package controllers

import (
	"blog/app/models"
	"blog/app/routes"
	"fmt"
	"github.com/revel/revel"
)

type Sessions struct {
	Common
}

func (s Sessions) Register() revel.Result {
	return s.Render()
}

func (s Sessions) SignIn() revel.Result {
	return s.Render()
}

func (s Sessions) CreateUser(user models.User) revel.Result {
	if err := models.CreateRecord(&user); err != nil {
		s.Flash.Error(err.Error())
		return s.Redirect(routes.Sessions.Register())
	}
	s.Flash.Success(fmt.Sprintf("Welcome, %v %v",
		user.FirstName, user.LastName))
	s.Session["user"] = user.Email
	return s.Redirect(routes.Posts.Index())
}

func (s Sessions) FindUser(email string) revel.Result {
	var user models.User
	attrs := map[string]interface{}{"email": email}

	if err := models.FindByMap(attrs, &user, true); err != nil {
		s.Flash.Error(err.Error())
		return s.Redirect(routes.Sessions.SignIn())
	}
	s.Session["user"] = email
	return s.Redirect(routes.Posts.Index())
}
