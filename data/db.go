package data

import (
	"database/sql"
	"eddy.com/todo/entity"
	_ "github.com/lib/pq"
	"log"
	"os"
)

func getConection() (*sql.DB, error){
	return sql.Open("postgres", os.Getenv("DATABASE_URL"))
}

func InitDatabase() {

	log.Print("db file created")

	dataBase, err := getConection()
	if err != nil {
		log.Fatal(err.Error())
	}
	defer dataBase.Close()

	createTodoTable(dataBase)
}

func createTodoTable(db *sql.DB) {
	query := `CREATE TABLE IF NOT EXISTS todo (
		id SERIAL PRIMARY KEY,		
		description VARCHAR (50),
		finished BOOLEAN	
	  );`

	statemant, err := db.Prepare(query)
	if err != nil {
		log.Fatal(err.Error())
	}
	statemant.Exec()
	log.Print("todo table created")
}

func Insert(todo entity.Todo) error {
	db, _ := getConection()
	defer db.Close()

	const query = `INSERT INTO todo (description, finished) VALUES ($1,$2);`
	statemant, err := db.Prepare(query)
	if err != nil {
		log.Fatal(err.Error())
		return err
	}

	_, err = statemant.Exec(todo.Description, todo.Finished)
	if err != nil {
		log.Fatal(err.Error())
		return err
	}
	return nil
}

func FindAll() ([]entity.Todo, error) {
	db, _ := getConection()
	defer db.Close()

	var todoList []entity.Todo
	row, err := db.Query("SELECT * FROM todo")
	if err != nil {
		log.Fatal(err.Error())
		return nil, err
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
	return todoList, nil
}

func Update(todo entity.Todo) error {
	db, _ := getConection()
	defer db.Close()

	const query = `
		UPDATE todo 
			SET description = $1, 	
				finished = $2 
		WHERE id = $3;`
	statemant, err := db.Prepare(query)
	if err != nil {
		log.Fatal(err.Error())
		return err
	}

	_, err = statemant.Exec(
		todo.Description,
		todo.Finished,
		todo.Id,
		)
	if err != nil {
		log.Fatal(err.Error())
		return err
	}
	return nil
}

func Delete(todoId int) error {
	db, _ := getConection()
	defer db.Close()

	const query = `DELETE FROM todo WHERE id = $1;`
	statemant, err := db.Prepare(query)
	if err != nil {
		log.Fatal(err.Error())
		return err
	}

	_, err = statemant.Exec(todoId)
	if err != nil {
		log.Fatal(err.Error())
		return err
	}
	return nil
}

func SetToFinished(todoId int) error {
	db, _ := getConection()
	defer db.Close()

	const query = `
		UPDATE todo 
			SET finished = true
		WHERE id = $1;`
	statemant, err := db.Prepare(query)
	if err != nil {
		log.Fatal(err.Error())
		return err
	}

	_, err = statemant.Exec(todoId)
	if err != nil {
		log.Fatal(err.Error())
		return err
	}
	return nil
}