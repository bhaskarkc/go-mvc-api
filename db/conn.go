package db

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
)

var Db *sqlx.DB

func DbConnect() *sqlx.DB {
	var err error
	Db, err = sqlx.Open("mysql", "root:docker_mysql_pass@tcp(127.0.0.1:9906)/dn-uat")

	if err != nil {
		panic(err.Error())
	}

	log.Println("Database connection successful.")
	return Db
}
