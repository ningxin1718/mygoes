//@Time  : 2018/4/3 17:40
//@Author: Greg Li
package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

const (
	DB_USER     = "ramesh"
	DB_PASSWORD = "secret"
	DB_NAME     = "test_db"
)

func main() {
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		DB_USER, DB_PASSWORD, DB_NAME)
	db, err := sql.Open("postgres", dbinfo)
	checkErr(err)
	defer db.Close()

	rows, err := db.Query("SELECT COUNT(*) as count FROM  table_name")
	fmt.Println("Total count:",checkCount(rows))
	checkErr(err)
}

func checkCount(rows *sql.Rows) (count int) {
	for rows.Next() {
		err:= rows.Scan(&count)
		checkErr(err)
	}
	return count
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}