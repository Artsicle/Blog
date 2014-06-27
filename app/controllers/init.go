package controllers

import "github.com/revel/revel"

func init() {
	revel.InterceptMethod(App.AddUser, revel.BEFORE)
}
