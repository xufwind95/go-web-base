package index

import (
	"github.com/gin-gonic/gin"

	"github.com/xufwind95/go-web-base/controller"
	"github.com/xufwind95/go-web-base/service/index"
)

func CreateAppUser(c *gin.Context) {
	var userService index.UserService
	if err := c.Bind(&userService); err != nil {
		controller.SendResponse(c, err, nil)
		return
	}

	if err := userService.AddUser(); err != nil {
		controller.SendResponse(c, err, nil)
		return
	}

	controller.SendResponse(c, nil, "success!")
}

func GetUser(c *gin.Context) {
	user, err := index.FindUser("1")
	if err != nil {
		controller.SendResponse(c, err, nil)
		return
	}
	controller.SendResponse(c, nil, user)
}
