package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	_ "loveHome/models"
	_ "loveHome/routers"
	"net/http"
	"strings"
)

func main() {
	ignoreStaticPath()
	//models.TestUploadByFilename("main.go")
	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.Run(":8899") //优先使用此端口
}

func ignoreStaticPath() {

	//透明static
	beego.SetStaticPath("group1/M00/", "fdfs/storage_data/data/")

	beego.InsertFilter("/", beego.BeforeRouter, TransparentStatic)
	beego.InsertFilter("/*", beego.BeforeRouter, TransparentStatic)
}

func TransparentStatic(ctx *context.Context) {
	orpath := ctx.Request.URL.Path
	beego.Debug("request url: ", orpath)
	//如果请求uri还有api字段,说明是指令应该取消静态资源路径重定向
	if strings.Index(orpath, "api") >= 0 {
		return
	}
	http.ServeFile(ctx.ResponseWriter, ctx.Request, "static/html/"+ctx.Request.URL.Path)
}
