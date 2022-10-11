package adm

import (
	"time"

	"svr/app/auth/jwt"
	"svr/app/auth/store"
	"svr/app/util"
	"svr/app/util/logx"
	"svr/app/util/rdx"

	"svr/app/cfg"
	"svr/app/db"

	"github.com/gin-gonic/gin"
)

func Run(file string) {
	cfg.InitBase(file)
	cfg.InitLog("appname")
	// cfg.InitCache()
	cfg.InitSMS()
	cfg.InitOss()
	cfg.InitDB()
	db.InitDB(cfg.DB)
	err := rdx.InitRdb(cfg.Cache.Addr, cfg.Cache.Password, cfg.Cache.Database)
	if err != nil {
		logx.Fatal(err)
	}

	// 是否打印访问日志, 在非正式环境都打印
	// if mode != "release" {
	// 	// engine.Use(gin.Logger())
	// }
	if cfg.DebugCurl {
		util.CurlDebug = true
	}
	store, err := store.NewRedis(cfg.Cache.Addr, cfg.Cache.Password, cfg.Cache.Database)
	if err != nil {
		logx.Fatal(err)
	}
	authManager := jwt.NewManager(store, time.Duration(cfg.Auth.TokenValidMinutes)*time.Minute, cfg.Auth.TokenSignSecret)
	// debug, release, test
	mode := "release"
	if cfg.Debug {
		mode = "debug"
	}
	gin.SetMode(mode)
	engine := gin.Default()
	server := apiserver.NewAPIServer(authManager)
	server.InitRouter(engine)
	engine.Run(cfg.PortAPI)
	logx.Info("Server exiting")
}
