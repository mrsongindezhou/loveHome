package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"loveHome/models"
)

type AreaController struct {
	beego.Controller
}

func (c *AreaController) RetData(resp map[string]interface{}) {
	c.Data["json"] = &resp
	c.ServeJSON()
}

func (c *AreaController) GetArea() {
	beego.Info("connect successful")
	resp := make(map[string]interface{})
	defer c.RetData(resp)
	//从session获取数据

	//从数据库获取数据
	area := []models.Area{} //定义一个接收数据对对象
	o := orm.NewOrm()       //获取连接
	n, err := o.QueryTable("area").All(&area)
	if err != nil {
		beego.Info("o.Read err:", err)
		resp["errno"] = 4001
		resp["errmsg"] = "查询失败"
		return
	}
	if n == 0 {
		resp["errno"] = 4002
		resp["errmsg"] = "没有查到数据"
	}
	resp["errno"] = 0
	resp["errmsg"] = "OK"
	resp["data"] = &area
	//打包成json格式，返回给客户端

}
