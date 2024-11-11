package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
)

// type MainController struct {
// 	beego.Controller
// }

type (
	MainController struct {
		BaseController
	}
)

func (c *MainController) Get() {
	beego.BConfig.WebConfig.AutoRender = true
	c.PublicContent("index")
	c.Data["HTMLTitle"] = "Books Rest API"
}

// func (c *MainController) Get() {
// 	c.Data["Website"] = "beego.me"
// 	c.Data["Email"] = "astaxie@gmail.com"
// 	c.TplName = "index.tpl"
// }
