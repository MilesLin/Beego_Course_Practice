package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

//export GetMethod
func (c *MainController) Get() {
	c.TplName = "home.tpl"
}
