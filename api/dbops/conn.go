package dbops

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var (
	dbConn *sql.DB
	err    error
)

// 初始化连接数据库
func init() {
	dbConn, err = sql.Open("mysql", "root:root@tcp(192.168.0.151:3306)/stream_media?charset=utf8")
	if err != nil {
		panic(err.Error())
	}
}
