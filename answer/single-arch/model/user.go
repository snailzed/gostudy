package model

import (
	"gostudy/answer/single-arch/id_gen"
	"gostudy/answer/single-arch/util"
	"time"
)

type UserInfo struct {
	Id       int64  `json:"user_id" db:"id"`
	Nickname string `json:"nickname" db:"nickname"`
	Sex      int    `json:"sex" db:"sex"`
	Salt     string `json:"_" db:"salt"`
	Username string `json:"user" db:"username"`
	Email    string `json:"email" db:"email"`
	Password string `json:"password" db:"password"`
}

func (u *UserInfo) GetUserByEmail() (user UserInfo, err error) {
	sql := "SELECT id,email,salt,password FROM user WHERE email = ?"
	stmt, err := Db.Preparex(sql)
	if err != nil {
		return user, err
	}
	defer stmt.Close()
	err = stmt.Get(&user, u.Email)
	return
}

func (u *UserInfo) Add() (err error) {
	sql := "INSERT INTO user (id,username,email,salt,password,create_time) values(?,?,?,?,?,?)"
	stmt, err := Db.Preparex(sql)
	if err != nil {
		return err
	}
	defer stmt.Close()
	salt := util.RandString(16)
	password := util.Md5([]byte(u.Password + salt))
	id, err := id_gen.GetId()
	if err != nil {
		return err
	}
	_, err = stmt.Exec(id, u.Username, u.Email, salt, password, time.Now().Unix())
	if err != nil {
		return err
	}
	return nil
}
