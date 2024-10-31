package mysql

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func init() {
	// 连接数据库
	tmpDB, err := sql.Open("mysql", "user:password@tcp(localhost:3306)/database_name")
	if err != nil {
		panic(err.Error())
	}

	// 测试连接
	err = tmpDB.Ping()
	if err != nil {
		panic(err.Error())
	}
	db = tmpDB
}
