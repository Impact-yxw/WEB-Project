package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
{
	"code": 1001, // 程序中的错误码
	"msg": xx,    // 提示信息
	"data": xxx,  // 携带的数据
}

*/

type ResponseData struct {
	Code ResCode     `json:"code"`
	Msg  interface{} `json:"msg"`
	Data interface{} `json:"data"`
}

func ResponseError(c *gin.Context, code ResCode) {
	rd := &ResponseData{
		Code: code,
		Msg:  code.Msg(),
		Data: nil,
	}
	c.JSON(http.StatusOK, rd)

}

func ResponseErrorWithMsg(c *gin.Context, code ResCode, msg interface{}) {
	rd := &ResponseData{
		Code: code,
		Msg:  msg,
		Data: nil,
	}
	c.JSON(http.StatusOK, rd)

}

func ResponseSuccess(c *gin.Context, data interface{}) {
	rd := &ResponseData{
		Code: CodeSuccess,
		Msg:  CodeSuccess.Msg(),
		Data: data,
	}
	c.JSON(http.StatusOK, rd)

}

func ResponseNotFound(c *gin.Context) {
	rd := &ResponseData{
		Code: CodeNotFound,
		Msg:  CodeNotFound.Msg(),
		Data: nil,
	}
	c.JSON(http.StatusOK, rd)

}
