package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.Data["content"] = "./app/app.html"
	c.TplName = "starter.html"
}

func (c *MainController) Index() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.Data["content"] = "./app/app.html"
	c.TplName = "app/app.html"
}

func (c *MainController) Images() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.Data["content"] = "./images/images.html"
	c.TplName = "images/images.html"
}

func (c *MainController) Dashboard() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.Data["content"] = "./dashboard/dashboard.html"
	c.TplName = "dashboard/dashboard.html"
}

func (c *MainController) Login() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.Data["content"] = "./login.html"
	c.TplName = "login.html"
}

func (c *MainController) Product() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.Data["content"] = "./app/product.html"
	c.TplName = "app/product.html"
}

func (c *MainController) Container() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.Data["content"] = "./app/container.html"
	c.TplName = "app/container.html"
}

func (c *MainController) Configs() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.Data["content"] = "./app/configs.html"
	c.TplName = "app/configs.html"
}