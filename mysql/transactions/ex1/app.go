package main

import (
	"context"
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

	var sql string = "SELECT id, name, password FROM users"
	res, err := db.Query(sql)

	if err != nil {
		log.Fatal(err)
	}
	var user = Users{}
	var users = []Users{}
	for res.Next() {

		err := res.Scan(&user.Id, &user.Name, &user.Password)
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, user)
	}

	fmt.Println(users)

	//transaction

	ctx := context.Background()
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	_, err = tx.ExecContext(ctx, "INSERT INTO users (name, email, password) VALUES ('sa', 'sa@gmail.com', '123)'")
	if err != nil {
		tx.Rollback()
		return
	}

	// The next query is handled similarly
	_, err = tx.ExecContext(ctx, "INSERT INTO users (name, email, password) VALUES ('sa', 'sa@gmail.com', ''404)")
	if err != nil {
		tx.Rollback()
		return
	}

	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}

}
