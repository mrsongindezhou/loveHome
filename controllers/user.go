package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"loveHome/models"
)

type UserController struct {
	beego.Controller
}

func (c *UserController) Reg() {
	beego.Info("api/v1.0/users connect successful")
	resp := make(map[string]interface{})
	defer c.RetData(resp)
	//获取前端json数据
	json.Unmarshal(c.Ctx.Input.RequestBody, &resp)
	//插入数据库
	o := orm.NewOrm()
	user := models.User{}
	user.Password_hash = resp["password"].(string) //接口类型，需要断言
	user.Mobile = resp["mobile"].(string)
	user.Name = resp["mobile"].(string)
	id, err := o.Insert(&user)
	if err != nil {
		resp["errno"] = 400
		resp["errmsg"] = "注册失败"
		return
	}
	beego.Info("reg successful id = ", id)
	c.SetSession("name", user.Name)
	resp["errno"] = 0
	resp["errmsg"] = "OK"
}

func (c *UserController) RetData(resp map[string]interface{}) {
	c.Data["json"] = &resp
	c.ServeJSON()
}
