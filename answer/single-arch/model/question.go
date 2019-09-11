package model

import (
	"time"
)

//model

type Question struct {
	QuestionId    int64     `json:"question_id_number" db:"id"`
	Caption       string    `json:"caption" db:"caption"`
	Content       string    `json:"content" db:"content"`
	AuthorId      int64     `json:"author_id_number" db:"author_id"`
	CategoryId    int64     `json:"category_id" db:"category_id"`
	Status        int32     `json:"status" db:"status"`
	CreateTime    time.Time `json:"-" db:"create_time"`
	CreateTimeStr string    `json:"create_time"`
	QuestionIdStr string    `json:"question_id"`
	AuthorIdStr   string    `json:"author_id"`
}

type ApiQuestion struct {
	Question
	AuthorName string `json:"author_name"`
}

type ApiQuestionDetail struct {
	Question
	AuthorName   string `json:"author_name"`
	CategoryName string `json:"category_name"`
}

func (q *Question) AddQuestion() (err error) {
	sqlStr := "INSERT INTO question(id,caption,content,author_id,category_id,status,create_time) values(?,?,?,?,?,?,?)"
	stmt, err := Db.Preparex(sqlStr)
	if err != nil {
		return err
	}
	defer stmt.Close()
	//
	_, err = stmt.Exec(q.QuestionId, q.Caption, q.Content, q.AuthorId, q.CategoryId, q.Status, time.Now().Unix())
	if err != nil {
		return err
	}
	return nil
}
