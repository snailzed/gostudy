package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gostudy/answer/single-arch/filter"
	"gostudy/answer/single-arch/id_gen"
	"gostudy/answer/single-arch/middleware/account"
	"gostudy/answer/single-arch/model"
	"gostudy/answer/single-arch/util"
	"log"
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
	if len(question.Caption) == 0 || len(question.Content) == 0 || question.CategoryId == 0 {
		util.ResponseError(ctx, util.ErrEmptyParam)
		return
	}
	_, hit := filter.Replace(question.Caption, "***")
	if hit {
		util.ResponseError(ctx, util.ErrTitleContainWord)
		return
	}
	_, hit = filter.Replace(question.Content, "***")
	if hit {
		util.ResponseError(ctx, util.ErrContentContainWord)
		return
	}
	qid, err := id_gen.GetId()
	if err != nil {
		util.ResponseError(ctx, util.ErrServer)
		return
	}
	question.QuestionId = int64(qid)
	question.AuthorId, err = account.GetUserId(ctx)
	if err != nil {
		log.Println(err)
		util.ResponseError(ctx, util.ErrNeedUserLogin)
		return
	}
	err = question.AddQuestion()
	if err != nil {
		util.ResponseError(ctx, util.ErrServer)
		return
	}
	util.ResponseSuccess(ctx, question)
	return
}
