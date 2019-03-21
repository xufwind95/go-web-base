package main

import (
	"github.com/cihub/seelog"
	"github.com/gin-gonic/gin"
	"github.com/xufwind95/go-web-base/config"
	"github.com/xufwind95/go-web-base/model"
	"github.com/xufwind95/go-web-base/pkg/database"
	"github.com/xufwind95/go-web-base/pkg/log"
	"github.com/xufwind95/go-web-base/pkg/redis"
	"github.com/xufwind95/go-web-base/pkg/util"
	"github.com/xufwind95/go-web-base/router"
	"net/http"
	"time"
)

func main() {
	// 初始化日志
	if err := log.InitLogger(); err != nil {
		panic(err)
	}
	defer seelog.Flush()

	// 读取配置文件
	conf := config.Read()

	// 初始化数据库
	database.InitDB(conf)
	defer database.CloseDB()
	model.Migrate(database.DB)

	redis.InitRedisPool(conf)
	defer redis.CloseRedisPool()

	// 初始化jwt
	util.InitJwtConfig(conf)

	// 设置运行模式
	gin.SetMode(conf.RunModel)

	// 加载路由并启动服务
	g := gin.New()
	router.Load(g, conf)
	s := &http.Server{
		Addr:              conf.Port,
		Handler:           g,
		ReadHeaderTimeout: 10 * time.Second,
		WriteTimeout:      10 * time.Second,
		MaxHeaderBytes:    1 << 20,
	}
	if err := s.ListenAndServe(); err != nil {
		panic(err)
	}

	// 监听停止信号，web服务以goroutine的形式启动后，才需要这个中方式来阻塞主线程
	// c := make(chan os.Signal, 1)
	// signal.Notify(c, os.Interrupt)
	// signal.Notify(c, syscall.SIGTERM, syscall.SIGINT)
	// <-c
}
