package controllers

import "github.com/xxdu521/cmdbgo/server/controllers/auth"

type LayoutController struct {
	auth.LoginRequiredController
}

func (c *LayoutController) Prepare(){
	c.LoginRequiredController.Prepare()

	c.Layout = "layouts/base.html"
	c.LayoutSections = map[string]string{
		"LayoutStyle":"",
		"LayoutScript":"",
	}
	c.Data["menu"] = ""
	c.Data["expand"] = ""
}