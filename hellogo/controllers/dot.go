package controllers

import (
	. "github.com/astaxie/beego"
)

type DotController struct {
	Controller
}

func (c *DotController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}
