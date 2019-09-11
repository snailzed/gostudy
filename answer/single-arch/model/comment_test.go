package model

import (
	"gostudy/answer/single-arch/id_gen"
	"testing"
	"time"
)

func BenchmarkComment_AddComment(b *testing.B) {
	id_gen.Init(1)
	Init()
	for i := 0; i < b.N; i++ {
		id, err := id_gen.GetId()
		if err != nil {
			b.Errorf("Get Id error:%v\n", err)
		}
		comment := &Comment{
			Id:        int64(id),
			AnswerId:  1,
			AuthorId:  1,
			Content:   "当前Nano：" + string(time.Now().Format("2006-01-02 15:04:05")),
			VoteCount: 0,
		}
		err = comment.AddComment(0)
		if err != nil {
			b.Errorf("INsert error:%v\n", err)
		}
	}
}
