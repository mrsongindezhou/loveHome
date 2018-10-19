package controllers

import (
	"github.com/astaxie/beego"
	"loveHome/models"
)

type SessionController struct {
	beego.Controller
}

func (c *SessionController) GetSessionData() {
	beego.Info("api/v1.0/session connect successful")
	resp := make(map[string]interface{})
	defer c.RetData(resp)

	user := models.User{}
	resp["errno"] = models.RECODE_DBERR
	resp["errmsg"] = models.RecodeText(models.RECODE_DBERR)
	name := c.GetSession("name")
	if name != nil {
		user.Name = name.(string)
		resp["errno"] = models.RECODE_OK
		resp["errmsg"] = models.RecodeText(models.RECODE_OK)
		resp["data"] = user
	}
}

func (c *SessionController) RetData(resp map[string]interface{}) {
	c.Data["json"] = &resp
	c.ServeJSON()
}
