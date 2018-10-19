package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/cache"
	_ "github.com/astaxie/beego/cache/redis"
	"github.com/astaxie/beego/orm"
	"loveHome/models"
	"time"
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
	//打包成json格式，返回给客户端
	defer c.RetData(resp)
	//从redis缓存中读取数据
	bm, errConn := cache.NewCache("redis", `{"key":"lovehome","conn":"47.94.213.158:6379","dbNum":"0"}`)
	if errConn != nil {
		beego.Info("cache.NewCache err:", errConn)
	}
	cacheErr := bm.Put("aaa", "bbb", 10*time.Second)
	if cacheErr != nil {
		beego.Error("bm.Put err:", cacheErr)
	}
	//从mysql数据库获取数据
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
		return
	}
	resp["errno"] = 0
	resp["errmsg"] = "OK"
	resp["data"] = &area
}
