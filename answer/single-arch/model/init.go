package model

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var Db *sqlx.DB

//Db init
func Init() {
	var err error
	Db, err = sqlx.Connect("mysql", "root:root@tcp(127.0.0.1:3306)/mercury")
	if err != nil {
		panic(err)
	}
	err = Db.Ping()
	if err != nil {
		panic(err)
	}
}
