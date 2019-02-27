package tools

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/edumar111/fastpv-auth/helper"
)

//DBConn connect to db
func DBConn() (db *sql.DB) {
	dbDriver := "mysql"

	db, err := sql.Open(dbDriver, helper.MysqlHost)
	if err != nil {
		fmt.Println("error al conectar con la bd")
		panic(err.Error())
	}
	fmt.Println("connect db success ")
	return db
}