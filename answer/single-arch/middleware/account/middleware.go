package account

import (
	"github.com/gin-gonic/gin"
	"gostudy/answer/single-arch/util"
)

func AuthMiddleware(ctx *gin.Context) {
	ProcessRequest(ctx)
	if !IsLogin(ctx) {
		util.ResponseError(ctx, util.ErrNeedUserLogin)
		//中断请求
		ctx.Abort()
		return
	}
	//继续处理
	ctx.Next()
}
