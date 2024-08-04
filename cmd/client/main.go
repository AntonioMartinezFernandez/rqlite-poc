package main

import (
	"database/sql"

	_ "github.com/rqlite/gorqlite/stdlib"
)

func main() {
	db, err := sql.Open("rqlite", "http://localhost:4001/")
	if err != nil {
		panic(err)
	}
	_, execErr := db.Exec("CREATE TABLE users (id INTEGER, name TEXT)")
	if execErr != nil {
		panic(execErr)
	}
}
