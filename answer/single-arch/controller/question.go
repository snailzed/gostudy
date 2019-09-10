package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gostudy/answer/single-arch/model"
	"gostudy/answer/single-arch/util"
)

type QuestionController struct {
}

func (q *QuestionController) AddQuestion(ctx *gin.Context) {
	var question model.Question
	err := ctx.BindJSON(&question)
	if err != nil {
		fmt.Println(err)
		util.ResponseError(ctx, util.ErrEmptyParam)
		return
	}
	util.ResponseSuccess(ctx, question)
	return
}
