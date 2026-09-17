package main

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "taskManager.db")
	handleError(err)
	tm := &taskManager{db}

	defer func() { handleError(db.Close()) }()
	tm.reset()

	tm.addTask("thing", "do a thing")
	handleError(err)

	// var thing string
	// result.Next()
	// err = result.Scan(&thing)
	// fmt.Println(thing)
	// err = result.Close()
	// handleError(err)

	a := 1
	a += 1
}
