package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"

	"fmt"
	"log"
)

type Users struct {
	Id       int
	Name     string
	Password string
}

func main() {

	db, err := sql.Open("mysql", "sa:123@tcp(127.0.0.1:3306)/api-go")
	if err != nil {
		log.Fatal(err)
	}
	if db != nil {
		fmt.Println("Database connected successfully!!")
	}

	handleError(err)

	defer db.Close()

	tx, err := db.Begin()
	handleError(err)

	// insert a record into table1
	res, err := tx.Exec("INSERT INTO users (name, email, password) VALUES ('sa', 'sa@gmail.com', '123')")
	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}

	// fetch the auto incremented id
	id, err := res.LastInsertId()
	handleError(err)

	// insert record into table2, referencing the first record from table1
	res, err = tx.Exec("INSERT INTO users (name, email, password) VALUES (?,?,?)", id, "sa", "sa")
	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}

	// commit the transaction
	handleError(tx.Commit())

	log.Println("Done.")
}

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
