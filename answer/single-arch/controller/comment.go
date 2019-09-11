package controller

import (
	"github.com/gin-gonic/gin"
	"gostudy/answer/single-arch/model"
	"gostudy/answer/single-arch/util"
	"log"
	"strconv"
)

type CommentController struct {
}

//获取顶级评论
func (c *CommentController) RootComments(ctx *gin.Context) {
	pageStr := ctx.DefaultQuery("page", "1")
	limitStr := ctx.DefaultQuery("limit", "20")
	answerIdStr, exists := ctx.GetQuery("answer_id")
	if !exists {
		util.ResponseError(ctx, util.ErrEmptyParam)
		return
	}

	page, err := strconv.ParseInt(pageStr, 10, 64)
	if err != nil {
		util.ResponseError(ctx, util.ErrEmptyParam)
		return
	}
	limit, err := strconv.ParseInt(limitStr, 10, 64)
	if err != nil {
		util.ResponseError(ctx, util.ErrEmptyParam)
		return
	}
	answerId, err := strconv.ParseInt(answerIdStr, 10, 64)
	if err != nil {
		util.ResponseError(ctx, util.ErrEmptyParam)
		return
	}
	comments, err := model.GetComments(answerId, page, limit)
	if err != nil {
		log.Println(err)
		util.ResponseError(ctx, util.ErrServer)
		return
	}
	util.ResponseSuccess(ctx, comments)
}

//获取二级评论
func (c *CommentController) ChildComments(ctx *gin.Context) {
	pageStr := ctx.DefaultQuery("page", "1")
	limitStr := ctx.DefaultQuery("limit", "20")
	parentIdStr, exists := ctx.GetQuery("parent_id")
	if !exists {
		util.ResponseError(ctx, util.ErrEmptyParam)
		return
	}

	page, err := strconv.ParseInt(pageStr, 10, 64)
	if err != nil {
		util.ResponseError(ctx, util.ErrEmptyParam)
		return
	}
	limit, err := strconv.ParseInt(limitStr, 10, 64)
	if err != nil {
		util.ResponseError(ctx, util.ErrEmptyParam)
		return
	}
	parentId, err := strconv.ParseInt(parentIdStr, 10, 64)
	if err != nil {
		util.ResponseError(ctx, util.ErrEmptyParam)
		return
	}
	comments, err := model.GetChildComments(parentId, page, limit)
	if err != nil {
		log.Println(err)
		util.ResponseError(ctx, util.ErrServer)
		return
	}
	util.ResponseSuccess(ctx, comments)
}
