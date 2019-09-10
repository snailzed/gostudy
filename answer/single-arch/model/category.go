package model

import "database/sql"

type Category struct {
	CategoryId   string `json:"id" db:"id"`
	CategoryName string `json:"name" db:"name"`
}

//获取列表
func GetCategoryList() (categorys []*Category, err error) {
	sqlStr := "SELECT id,name from category"
	stmt, err := Db.Preparex(sqlStr)
	if err != nil {
		return
	}
	defer stmt.Close()
	err = stmt.Select(&categorys)
	if err != nil && err == sql.ErrNoRows {
		err = nil
		return
	}
	if err != nil {
		return nil, err
	}
	return
}
