package controllers

import "github.com/revel/revel"

func init() {
	revel.InterceptMethod(Common.AddUser, revel.BEFORE)
	revel.InterceptMethod(Posts.checkUser, revel.BEFORE)
}
