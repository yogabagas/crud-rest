package controllers

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type Context struct {
	DB *sql.DB
}

func DBintitial(*sql.DB) *sql.DB {
	var err error
	db, err := sql.Open("mysql", "root:root@/DBSelling")
	if err != nil {
		log.Fatal(err)
	}
	return db
}
