package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"loveHome/models"
)

type HouseIndexController struct {
	beego.Controller
}

func (c *HouseIndexController) RetData(resp map[string]interface{}) {
	c.Data["json"] = &resp
	c.ServeJSON()
}

func (c *HouseIndexController) GetHouseIndex() {
	beego.Info("api/v1.0/houses/index connect successful")
	resp := make(map[string]interface{})
	//打包成json格式，返回给客户端
	defer c.RetData(resp)
	houses := models.House{} //定义接收者
	o := orm.NewOrm()
	n, err := o.QueryTable("House").All(&houses)
	if err != nil {
		beego.Info("o.Read err:", err)
		resp["errno"] = 4001
		resp["errmsg"] = "查询失败"
		return
	}
	if n == 0 {
		resp["errno"] = 4002
		resp["errmsg"] = "没有查到数据"
		return
	}
	resp["errno"] = 0
	resp["errmsg"] = "OK"
	resp["data"] = &houses
}
