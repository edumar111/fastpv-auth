package tools

import (
	"database/sql"
	"fmt"
	"github.com/edumar111/fastpv-auth/helper"
)

func dbConn() (db *sql.DB) {
	dbDriver := "mysql"

	db, err := sql.Open(dbDriver, helper.MysqlHost)
	if err != nil {
		fmt.Println("error al conectar con la bd")
		panic(err.Error())
	}
	fmt.Println("connect db success ")
	return db
}