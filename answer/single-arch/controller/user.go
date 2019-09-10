package controller

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"gostudy/answer/single-arch/model"
	"gostudy/answer/single-arch/util"
	"log"
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
		log.Println(err)
		util.ResponseError(ctx, util.ErrServer)
		return
	}
	util.ResponseSuccess(ctx, nil)
	return
}

//登录
func (u *UserController) Login(ctx *gin.Context) {
	var userInfo model.UserInfo
	err := ctx.BindJSON(&userInfo)
	if err != nil {
		util.ResponseError(ctx, util.ErrEmptyParam)
		return
	}
	if len(userInfo.Email) == 0 || len(userInfo.Password) == 0 {
		util.ResponseError(ctx, util.ErrEmptyParam)
		return
	}
	dbUser, err := userInfo.GetUserByEmail()
	//记录为空
	if err == sql.ErrNoRows {
		util.ResponseError(ctx, util.ErrUserNotExists)
		return
	}
	if err != nil {
		util.ResponseError(ctx, util.ErrServer)
		return
	}
	//判断密码是否正确
	if dbUser.Password != util.Md5([]byte(userInfo.Password+dbUser.Salt)) {
		util.ResponseError(ctx, util.ErrUserOrPassword)
		return
	}
	util.ResponseSuccess(ctx, nil)
}
