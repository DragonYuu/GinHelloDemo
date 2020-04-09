package mysql

import (
	_"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var (
	DB *sqlx.DB
)

func GetDB() *sqlx.DB {
	return DB
}

func init() {
	//1. 连接数据库
	db, err := sqlx.Open("mysql", "root:ps@12345@tcp(172.81.250.202:3306)/sql_test?charset=utf8mb4")

	if err != nil {
		panic(err)
	}

	//defer db.Close()

	db.SetMaxIdleConns(20)
	db.SetMaxOpenConns(20)

	if err := db.Ping(); err != nil {
		panic(err)
	}
	DB = db
}
