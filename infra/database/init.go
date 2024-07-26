package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func Initialize() *sql.DB {
	dsn := "Ramit:fplOrZenc6BWjfmgiGKolM7OxYoMVFgr@tcp(svc-3482219c-a389-4079-b18b-d50662524e8a-shared-dml.aws-virginia-6.svc.singlestore.com:3333)/database_d5bb3?tls=skip-verify"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	fmt.Println("connected to vector db")
	return db
}
