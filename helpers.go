package main

import (
	"database/sql"
	"uuid"
)

type taskManager struct {
	db *sql.DB
}

// Add a task
func (t *taskManager) addTask(name string, description string) {
	_, err := t.db.Exec("insert into main.tasks values (?, ?, ?)", name, description, uuid.New().String())
	handleError(err)
}

func (t *taskManager) makeTaskTable() {
	query := `
create table if not exists tasks as
		select '' as name,
				'' as description,
				'' as uuid
		limit 0;
`
	_, err := t.db.Exec(query)
	handleError(err)
}
func (t *taskManager) reset() {
	_, err := t.db.Exec("drop table if exists main.tasks; ")
	handleError(err)
	t.makeTaskTable()
}

// type task struct {
// 	name string
// }

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}
