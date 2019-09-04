package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

//打开连接： sqlx.Open()
//查询：sqlx.DB.Get  sqlx.DB.Select
//插入修改删除： sqlx.DB.Exec
//事务： sqlx.DB.Begin()   sqlx.DB.Commit()  sqlx.DB.Rollback()
var Db *sqlx.DB

func initDb() (err error) {
	Db, err = sqlx.Open("mysql", "root:root@tcp(127.0.0.1:3306)/test")
	if err != nil {
		return err
	}
	Db.SetMaxOpenConns(100) //设置打开的最大连接数
	Db.SetMaxIdleConns(10)  //最大等待使用的连接数
	return nil
}

type User struct {
	Id   int    `db:"id"`
	Name string `db:"name"`
	Age  int    `db:"age"`
}

func testQuery() (err error) {
	sqlStr := "select id,name,age from user where id = ?"
	var user User
	err = Db.Get(&user, sqlStr, 2)
	if err != nil && err != sql.ErrNoRows {
		return err
	}
	fmt.Printf("user:%+v\n", user)
	return nil
}

//预处理
// 预处理是为了提高查询性能；
//// 实现的原理是先将查询语句发送给Mysql数据库做预解析；
//// 然后再将需要查询的条件数据发送给Mysql数据库进行执行；
//// 这种原理类似于GO语言和Python语言执行效率的对比；
//// Go语言是需要先编译的，Python是一边执行一边编译。
func PrepareQuery(id int) {
	stmt, err := Db.Prepare("select id, name, age from user where id= ?")
	if err != nil {
		panic(err)
	}

	rows, err := stmt.Query(id)
	if err != nil {
		panic(err)
	}

	defer stmt.Close()
	defer rows.Close()

	for rows.Next() {
		var user User
		err := rows.Scan(&user.Id, &user.Name, &user.Age)
		if err != nil {
			panic(err)
		}
		fmt.Printf("user: %#v\n", user)
	}
}

//预处理
func PrepareQueryV2() {
	stmt, err := Db.Preparex("select id,name,age from user where id > ?")
	if err != nil {
		fmt.Println("Preparex failed:", err)
	}
	defer stmt.Close()
	var users []User
	err = stmt.Select(&users, 0)
	if err != nil {
		fmt.Println("stmt select failed.", err)
	}
	fmt.Printf("users:%#v\n", users)
}
func PrepareExec() {
	stmt, err := Db.Preparex("update user set name=?, age=? where id=?")
	if err != nil {
		fmt.Println("Preparex failed:", err)
	}
	defer stmt.Close()
	ret, err := stmt.Exec("Snailzed", 25, 1)

	fmt.Println(ret.RowsAffected())
}

//事务
func Transaction() {

	// 开启事务
	tx, err := Db.Begin()

	if err != nil {
		panic(err)
	}

	result, err := tx.Exec("insert into user(name, age)values(?,?)", "Jack", 98)
	if err != nil {
		// 失败回滚
		tx.Rollback()
		panic(err)
	}

	fmt.Println("result", result)

	exec, err := tx.Exec("update user set name=?, age=? where id=?", "Jack", 98, 1)
	if err != nil {
		// 失败回滚
		tx.Rollback()
		panic(err)
	}
	fmt.Println("exec", exec)

	// 提交事务
	err = tx.Commit()

	if err != nil {
		// 失败回滚
		tx.Rollback()
		panic(err)
	}
}

func Update() {
	name := "Miles"
	age := 88
	id := 3

	result, err := Db.Exec("update user set name=?, age=? where id=?", name, age, id)
	if err != nil {
		panic(err)
	}

	// RowsAffected returns the number of rows affected by an
	// update, insert, or delete.
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		panic(err)
	}

	fmt.Printf("update id:%d, affect rows:%d\n", id, rowsAffected)

}

func Insert() {
	name := "Lucy"
	age := 18

	result, err := Db.Exec("insert into user(name, age) values (?,?)", name, age)
	if err != nil {
		panic(err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}

	affected, err := result.RowsAffected()
	if err != nil {
		panic(err)
	}

	fmt.Printf("last insert id:%d affect rows:%d\n", id, affected)
}
func Run() {
	err := initDb()
	if err != nil {
		panic(err)
	}
	err = testQuery()
	if err != nil {
		panic(err)
	}
	Insert()
	Update()
	PrepareQuery(1)
	Transaction()
	PrepareQueryV2()
	PrepareExec()
}
