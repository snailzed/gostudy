package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gostudy/answer/single-arch/model"
	"gostudy/answer/single-arch/util"
)

type CategoryController struct {
}

//获取分类列表
func (c *CategoryController) GetCategoryList(ctx *gin.Context) {
	data, err := model.GetCategoryList()
	fmt.Println(data, err)
	if err != nil {
		util.ResponseError(ctx, util.ErrServer)
		return
	}
	util.ResponseSuccess(ctx, data)
}

func (c *CategoryController) GetQuestionList(ctx *gin.Context) {

}
