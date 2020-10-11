package models

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Task struct {
	ID          string
	Description string
	Deadline    string
	Priority    string
}

var Db *sql.DB

func Init() {
	var err error
	Db, err = sql.Open("postgres", "postgres://jyfhuang:password@localhost/myworld?sslmode=disable")
	if err != nil {
		panic(err)
	}

	if err = Db.Ping(); err != nil {
		panic(err)
	}
	fmt.Println("You connected to your database.")

}
