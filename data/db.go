package data

import (
	"database/sql"
	"log"
	"os"

	"eddy.com/todo/entity"
	_ "github.com/mattn/go-sqlite3"
)

func InitDatabase() {

	if _, err := os.Stat("database.db"); os.IsNotExist(err) {
		file, err := os.Create("database.db")
		if err != nil {
			log.Fatal(err.Error())
		}

		file.Close()

		log.Print("db file created")

		dataBase, _ := sql.Open("sqlite3", "./database.db")
		defer dataBase.Close()

		log.Print("db created")

		createTodoTable(dataBase)

		return
	}
	log.Print("db already exist")
}

func createTodoTable(db *sql.DB) {
	query := `CREATE TABLE todo (
		id integer NOT NULL PRIMARY KEY AUTOINCREMENT,		
		description TEXT,
		finished INT	
	  );`

	statemant, err := db.Prepare(query)
	if err != nil {
		log.Fatal(err.Error())
	}
	statemant.Exec()
	log.Print("todo table created")
}

func InsertTodo(todo entity.Todo) {
	db, _ := sql.Open("sqlite3", "./database.db")
	defer db.Close()

	query := `INSERT INTO todo (description, finished) VALUES (?,?)`
	statemant, err := db.Prepare(query)
	if err != nil {
		log.Fatal(err.Error())
	}

	_, err = statemant.Exec(todo.Description, todo.Finished)
	if err != nil {
		log.Fatal(err.Error())
	}
}

func FindTodo() []entity.Todo {
	db, _ := sql.Open("sqlite3", "./database.db")
	defer db.Close()

	var todoList []entity.Todo
	row, err := db.Query("SELECT * FROM todo")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer row.Close()

	for row.Next() {
		var todo entity.Todo

		row.Scan(
			&todo.Id,
			&todo.Description,
			&todo.Finished,
		)
		todoList = append(todoList, todo)
	}
	return todoList
}
