package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/xufwind95/go-web-base/pkg/app_error"
)

/*
	统一接口输出格式和输出函数
	control 调用 service ，service 调用 model ,control可使用service的数据结构，service可使用model的数据结构
*/

type AppResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func SendResponse(c *gin.Context, err error, data interface{}) {
	code, message := app_error.DecodeError(err)

	c.JSON(
		http.StatusOK,
		AppResponse{
			Code:    code,
			Message: message,
			Data:    data,
		},
	)
}
