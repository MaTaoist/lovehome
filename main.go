package main

import (
	_ "lovehome/routers"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	_ "lovehome/models"
	"github.com/astaxie/beego/context"
	"strings"
	"net/http"
)

func main() {
	ignoreStaticPath()
	beego.Run()//设置端口，conf也可以设置不过此处优先级高
}

func ignoreStaticPath() {

	//透明static

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