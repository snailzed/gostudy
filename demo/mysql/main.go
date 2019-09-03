package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

//使用默认的mysql驱动
func main() {
	//mysqlNative()
	Run()
}

func mysqlNative() {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/test")
	if err != nil {
		fmt.Printf("open error:%v\n", err)
	}
	defer db.Close()
	// Open doesn't open a connection. Validate DSN data:
	err = db.Ping()
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	row, err := db.Query("select id,name,age from user where id = ?", 2)
	if err != nil {
		panic(err)
	}
	var user User
	if row.Next() {
		err = row.Scan(&user.Id, &user.Name, &user.Age)
		if err != nil && err != sql.ErrNoRows {
			panic(err)
		}
	}
	fmt.Printf("user:%+v\n", user)
}
