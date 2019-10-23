package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func DbConnect() *sql.DB {
	db, err := sql.Open("mysql", "root:docker_mysql_pass@tcp(127.0.0.1:9906)/dn-uat")

	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("Database connection successful.")
	return db
}
