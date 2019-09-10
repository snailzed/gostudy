package util

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func ResponseSuccess(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, &Response{
		Code: SUCCESS,
		Msg:  "OP_OK",
		Data: data,
	})
}

func ResponseError(ctx *gin.Context, code int) {
	ctx.JSON(http.StatusOK, &Response{
		Code: code,
		Msg:  GetErrorMessage(code),
	})
}
