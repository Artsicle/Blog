package controllers

import (
	"blog/app/models"
	"blog/app/routes"
	"fmt"
	"github.com/revel/revel"
	"regexp"
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

func (s Sessions) CreateUser(user models.User, password, passwordConfirmation string) revel.Result {
	s.Validation.Required(user.Email)
	s.Validation.Required(password)
	s.Validation.Required(passwordConfirmation)
	s.Validation.MinSize(password, 8)
	s.Validation.Match(password, regexp.MustCompile(passwordConfirmation))

	if s.Validation.HasErrors() {
		s.Validation.Keep()
		s.FlashParams()
		return s.Redirect(routes.Sessions.Register())
	}

	if err := user.GenerateHashedPassword(password); err != nil {
		s.Flash.Error(err.Error())
		return s.Redirect(routes.Sessions.Register())
	}

	if err := models.CreateRecord(&user); err != nil {
		s.Flash.Error(err.Error())
		return s.Redirect(routes.Sessions.Register())
	}

	s.Flash.Success(fmt.Sprintf("Welcome, %v %v",
		user.FirstName, user.LastName))
	s.Session["user"] = user.Email
	return s.Redirect(routes.Posts.Index())
}

func (s Sessions) FindUser(email, password string) revel.Result {
	var user models.User
	attrs := map[string]interface{}{"email": email}

	if err := models.FindByMap(attrs, &user, true); err != nil {
		s.Flash.Error(err.Error())
		return s.Redirect(routes.Sessions.SignIn())
	}

	if valid := user.CheckPassword(password); !valid {
		s.Flash.Error("Incorrect Password")
		return s.Redirect(routes.Sessions.SignIn())
	}

	s.Session["user"] = email
	return s.Redirect(routes.Posts.Index())
}
