package service

import (
	"eddy.com/todo/data"
	"eddy.com/todo/entity"
)

var todoList []entity.Todo
var id int

func Insert(todo entity.Todo) {
	data.InsertTodo(todo)
}

func Find() []entity.Todo {
	return data.FindTodo()
}

func SetToFinished(todoId int) {
	for i, todo := range todoList {
		if todoId == todo.Id {
			todoList[i].Finished = true
		}
	}
}
