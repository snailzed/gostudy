package model

import (
	"database/sql"
	"fmt"
	"log"
	"time"
)

type Comment struct {
	Id                int64     `json:"id" db:"id"`
	AnswerId          int64     `json:"-" db:"answer_id"`
	AuthorId          int64     `json:"author_id" db:"author_id"`
	ReplyAuthorId     int64     `json:"-" db:"reply_author_id"`
	Content           string    `json:"content" db:"content"`
	VoteCount         int       `json:"vote_count" db:"vote_count"`
	ParentId          int64     `json:"-" db:"parent_id"`
	Featured          int       `json:"featured" db:"is_featured"`
	IsAuthor          int       `json:"is_author" db:"is_author"`
	ChildCommentCount int       `json:"child_comment_count" db:"child_comment_count"`
	IsDelete          int       `json:"is_delete" db:"is_delete"`
	AllowVote         int       `json:"allow_vote" db:"allow_vote"`
	AllowReply        int       `json:"allow_reply" db:"allow_reply"`
	CreateTime        int       `json:"create_time" db:"create_time"`
	Author            *UserInfo `json:"author"`
}

//顶级评论
type RootComment struct {
	ChildComments []*ChildComment `json:"child_comments"`
	Comment
}

//子评论
type ChildComment struct {
	Comment
	ReplyAuthor *UserInfo `json:"reply_to_author"`
}

//获取顶级评论，顶级评论没有被回复的用户信息
func GetComments(anwserId, page, limit int64) (comments []*RootComment, err error) {
	sqlStr := "SELECT id,answer_id,author_id,content,vote_count,is_author,is_author,child_comment_count,is_delete,allow_vote,allow_reply,create_time FROM comment WHERE parent_id = 0 AND answer_id = ? ORDER BY is_featured DESC,create_time ASC LIMIT ?,?"
	stmt, err := Db.Preparex(sqlStr)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	err = stmt.Select(&comments, anwserId, page-1, limit)
	if err != nil && err == sql.ErrNoRows {
		return comments, nil
	}
	if err != nil {
		return nil, err
	}
	for _, comment := range comments {
		//如果有子评论s
		if comment.ChildCommentCount > 0 {
			comment.ChildComments, err = GetChildComments(comment.Id, 1, 2)
			fmt.Println(comment.ChildComments)
			if err != nil {
				log.Println(err)
				continue
			}
		}
		//用户信息
		comment.Author, err = GetUserInfoById(comment.AuthorId)
		if err != nil {
			log.Println(err)
			continue
		}
	}
	return
}

//获取子评论,有被回复的用户信息
func GetChildComments(parent_id int64, page, limit int64) (comments []*ChildComment, err error) {
	sqlStr := "SELECT id,answer_id,author_id,content,vote_count,reply_author_id,is_author,is_author,child_comment_count,is_delete,allow_vote,allow_reply,create_time FROM comment WHERE parent_id = ? ORDER BY is_featured DESC,create_time ASC LIMIT ?,?"
	stmt, err := Db.Preparex(sqlStr)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	err = stmt.Select(&comments, parent_id, page-1, limit)
	if err != nil && err == sql.ErrNoRows {
		return comments, nil
	}
	if err != nil {
		return nil, err
	}
	for _, comment := range comments {
		//获取用户信息
		//用户信息
		comment.Author, err = GetUserInfoById(comment.AuthorId)
		if err != nil {
			log.Println(err)
			continue
		}
		comment.ReplyAuthor, err = GetUserInfoById(comment.ReplyAuthorId)
		if err != nil {
			log.Println(err)
			continue
		}
	}
	return
}

//添加评论
func (c *Comment) AddComment(parent_id int64) (err error) {
	sqlStr := "INSERT INTO comment(id,answer_id,author_id,content,parent_id,vote_count,create_time) values(?,?,?,?,?,?,?)"
	stmt, err := Db.Preparex(sqlStr)
	if err != nil {
		return err
	}
	defer stmt.Close()
	//
	_, err = stmt.Exec(c.Id, c.AnswerId, c.AuthorId, c.Content, parent_id, c.VoteCount, time.Now().Unix())
	if err != nil {
		return err
	}
	return nil
}
