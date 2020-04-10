package dao

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	sql "github.com/jmoiron/sqlx"
)

var db *sql.DB

func InitMySQL() (err error) {
	fmt.Println("InitMySQL....")
	if db == nil {
		db, err = sql.Connect("mysql", "root:ps@12345@tcp(172.81.250.202:3306)/gin_blog")
		if err != nil {
			return
		}
	}
	err = CreateTableWithUser() // 创建用户表
	if err != nil {
		return
	}
	err = CreateTableWithArticle() // 创建文章表
	if err != nil {
		return
	}
	return
}

//创建用户表
func CreateTableWithUser() (err error) {
	sqlStr := `CREATE TABLE IF NOT EXISTS users(
        id INT(4) PRIMARY KEY AUTO_INCREMENT NOT NULL,
        username VARCHAR(64),
        password VARCHAR(64),
        status INT(4),
        create_time INT(10)
        );`

	_, err = ModifyDB(sqlStr)
	return
}

// 创建文章表
func CreateTableWithArticle() (err error) {
	sqlStr := `create table if not exists article(
        id int(4) primary key auto_increment not null,
        title varchar(30),
        author varchar(20),
        tags varchar(30),
        short varchar(255),
        content longtext,
        create_time int(10),
        status int(4)
        );`
	_, err = ModifyDB(sqlStr)
	return
}

//操作数据库
func ModifyDB(sql string, args ...interface{}) (int64, error) {
	result, err := db.Exec(sql, args...)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	count, err := result.RowsAffected()
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	return count, nil
}

//查询
func QueryRowDB(dest interface{}, sql string, args ...interface{}) error {
	return db.Get(dest, sql, args...)
}

// 查询多条
func QueryRows(dest interface{}, sql string, args ...interface{}) error {
	return db.Select(dest, sql, args...)
}
