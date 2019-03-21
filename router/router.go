package router

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/xufwind95/go-web-base/config"
	"github.com/xufwind95/go-web-base/controller/index"
	"github.com/xufwind95/go-web-base/router/middleware"
)

func Load(g *gin.Engine, conf *config.Config) *gin.Engine {
	// 恢复之前调用失败的接口的功能
	g.Use(gin.Recovery())

	// 全局配置
	g.Use(func(c *gin.Context) {
		c.Set("conf", conf)
	})

	// 跨域
	g.Use(middleware.Cors())

	// 404 页面
	g.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "Incorrect API route.")
	})

	g.GET("", index.Welcome)
	g.GET("/test", index.Welcome2)

	u := g.Group("/user")
	u.Use(middleware.AuthMiddleware())
	{
		u.GET("/index", index.Welcome)
		u.POST("", index.CreateAppUser)
		u.GET("", index.GetUser)
	}

	return g
}
