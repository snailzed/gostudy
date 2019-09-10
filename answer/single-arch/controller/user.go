package controller

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"gostudy/answer/single-arch/model"
	"gostudy/answer/single-arch/util"
)

type UserController struct {
}

//注册
func (u *UserController) Register(ctx *gin.Context) {
	var userInfo model.UserInfo
	err := ctx.BindJSON(&userInfo)
	if err != nil {
		fmt.Println(err)
		util.ResponseError(ctx, util.ErrEmptyParam)
		return
	}
	if len(userInfo.Email) == 0 || len(userInfo.Password) == 0 || len(userInfo.Username) == 0 {
		util.ResponseError(ctx, util.ErrEmptyParam)
		return
	}

	//查找是否存在
	user, err := userInfo.GetUserByEmail()
	if err != nil && err != sql.ErrNoRows {
		util.ResponseError(ctx, util.ErrServer)
		return
	}
	if err == nil && user.Id > 0 {
		util.ResponseError(ctx, util.ErrUserExists)
		return
	}
	err = userInfo.Add()
	if err != nil {
		util.ResponseError(ctx, util.ErrServer)
		return
	}
	util.ResponseSuccess(ctx, userInfo)
	return
}

//登录
func (u *UserController) Login(ctx *gin.Context) {

}
