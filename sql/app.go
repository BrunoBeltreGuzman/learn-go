package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	var connection *sql.DB = getConnection()

	if connection != nil {
		fmt.Println("Database connection successfull")
	} else {
		fmt.Println(connection)
	}

}

func getConnection() (db *sql.DB) {

	const driver string = "mysql"
	const database string = "go"
	const user string = "sa"
	const password string = "admin"

	db, err := sql.Open(driver, user+":"+password+"@tcp(127.0.0.1:3306)/"+database)
	defer db.Close()

	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Database `" + database + "` connection successfull")
	}
	return db
}
