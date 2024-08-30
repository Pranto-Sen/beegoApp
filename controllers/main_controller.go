

package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	// Get the URL path
	path := c.Ctx.Input.URL()

	// Serve the same template for all the defined routes
	switch path {
	case "/", "/breed", "/favourite", "/voting":
		c.TplName = "index.tpl"
	default:
		// Handle 404 Not Found for undefined routes
		c.Abort("404")
	}
}



