package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"loveHome/models"
)

type SessionController struct {
	beego.Controller
}

/**
  获取session
*/
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

/**
  退出登录
*/
func (c *SessionController) DeleteSessionData() {
	resp := make(map[string]interface{})
	defer c.RetData(resp)
	c.DelSession("name")
	resp["errno"] = models.RECODE_OK
	resp["errmsg"] = models.RecodeText(models.RECODE_OK)
}

/**
  登录
*/
func (c *SessionController) Login() {
	resp := make(map[string]interface{})
	defer c.RetData(resp)
	//1.获取参数
	json.Unmarshal(c.Ctx.Input.RequestBody, &resp)
	beego.Info("入参：login info:", resp["mobile"], resp["password"])
	//2.判断合法性
	if resp["mobile"] == nil || resp["password"] == nil {
		resp["errno"] = models.RECODE_DATAERR
		resp["errmsg"] = models.RecodeText(models.RECODE_DATAERR)
		return
	}
	//3.连接数据库，查询用户
	user := models.User{}
	user.Mobile = resp["mobile"].(string)
	user.Password_hash = resp["password"].(string)
	o := orm.NewOrm()
	qt := o.QueryTable("User")
	err := qt.Filter("mobile", resp["mobile"].(string)).One(&user)
	if err != nil {
		resp["errno"] = models.RECODE_DATAERR
		resp["errmsg"] = models.RecodeText(models.RECODE_DATAERR)
		return
	}
	if user.Password_hash != resp["password"] {
		resp["errno"] = models.RECODE_DATAERR
		resp["errmsg"] = models.RecodeText(models.RECODE_DATAERR)
		return
	}
	//4.添加session
	c.SetSession("name", user.Name)
	c.SetSession("mobile", user.Mobile)
	c.SetSession("user_id", user.Id)
	//5.将json返回给前端
	resp["errno"] = models.RECODE_OK
	resp["errmsg"] = models.RecodeText(models.RECODE_OK)

}
func (c *SessionController) RetData(resp map[string]interface{}) {
	c.Data["json"] = &resp
	c.ServeJSON()
}
