package main

import (
	"flag"

	"github.com/gin-gonic/gin"

	"svr/app/util/logx"

	"svr/app/cfg"
	"svr/app/mw"
	"svr/cmd/api/hdl"
)

// 后台管理面板
// 前端项目在 https://git.xk.design/xkAdmin/slmAdm.git

func main() {
	file := flag.String("env", ".env", "env file path")
	flag.Parse()
	cfg.InitBase(*file)
	cfg.InitLog("appname")
	cfg.InitCache()
	cfg.InitOss()
	cfg.InitDB()
	r := route()
	if err := r.Run(cfg.PortAPI); err != nil {
		logx.Warn("Server run error : " + err.Error())
	}
	logx.Info("Server exiting")
}

func route() *gin.Engine {
	// debug, release, test
	mode := "release"
	if cfg.Debug {
		mode = "debug"
	}
	gin.SetMode(mode)
	r := gin.Default()
	r.Use(gin.Recovery())
	// 是否打印访问日志, 在非正式环境都打印
	if mode != "release" {
		r.Use(gin.Logger())
	}
	r.GET("/favicon.ico", func(c *gin.Context) { c.Status(204) })
	r.GET("/ping", func(c *gin.Context) { c.Writer.WriteString("pong") })

	r.POST("/sign/in", hdl.SignIn)

	r.Use(mw.Auth)
	r.GET("/sign/out", hdl.SignOut)

	r.GET("/info", hdl.Info)

	return r
}
