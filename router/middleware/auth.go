package middleware

import (
	"fmt"

	log "github.com/cihub/seelog"
	"github.com/gin-gonic/gin"

	"github.com/xufwind95/go-web-base/controller"
	"github.com/xufwind95/go-web-base/model"
	"github.com/xufwind95/go-web-base/pkg/app_error"
	"github.com/xufwind95/go-web-base/pkg/database"
	"github.com/xufwind95/go-web-base/pkg/util"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.Request.Header.Get("Authorization")

		if len(header) == 0 {
			controller.SendResponse(c, app_error.ErrTokenInvalid, nil)
			c.Abort()
			return
		}

		var t string
		// 从header中获取token
		fmt.Sscanf(header, "Bearer %s", &t)

		id, err := util.ParseToken(t)
		if err != nil {
			controller.SendResponse(c, app_error.ErrTokenInvalid, nil)
			c.Abort()
			return
		}

		// 将用户保存到上下文中
		user := model.UserTestModel{}
		if err := database.DB.Where("id = ?", id).First(&user).Error; err != nil {
			log.Error(fmt.Sprintf("not found user_id: %d", id))
			fmt.Println(err, "......")
			controller.SendResponse(c, app_error.ErrTokenInvalid, nil)
			c.Abort()
			return
		}
		c.Set("user", user)
		c.Next()
	}
}
