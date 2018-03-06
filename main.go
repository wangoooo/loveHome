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

	//开启session功能
	beego.BConfig.WebConfig.Session.SessionOn = true

	beego.SetStaticPath("/group1/M00", "data")

	/*	//测试fdfs接口
		_, _, err := models.FDFSUploadByFilename("./123.jpg")
		if err == nil {
		}
	*/
	beego.Run()
}

func ignoreStaticPath() {
	//透明static
	beego.InsertFilter("/", beego.BeforeRouter, TransparentStatic)
	beego.InsertFilter("/*", beego.BeforeRouter, TransparentStatic)
}

func TransparentStatic(ctx *context.Context) {
	orpath := ctx.Request.URL.Path

	beego.Debug("request url : ", orpath)

	//如果请求url还有api字段，说明是指令应该取消静态资源路径重定向
	if strings.Index(orpath, "api") >= 0 {
		return
	}

	//  /index.html ---> /static/html/index.html
	http.ServeFile(ctx.ResponseWriter, ctx.Request, "static/html/"+ctx.Request.URL.Path)

}
