package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	path := c.Ctx.Input.URL()

	switch path {
	case "/", "/breed", "/favourite", "/voting":
		c.TplName = "index.tpl"
	default:
		c.Abort("404")
	}
}



