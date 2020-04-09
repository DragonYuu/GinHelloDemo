package main

import (
	"GinHelloDemo/mysql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Id       int
	Username string
	Password string
}

func main() {
	db := mysql.GetDB()
	var user []User
	err := db.Select(&user, "SELECT id,username,password FROM users")
	if err != nil {
		fmt.Println(err)
	}
	for key, value := range user {
		fmt.Println(key, value.Id, value.Password, value.Username)
	}
	db.Close()
}
