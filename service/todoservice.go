package service

import (
	"eddy.com/todo/data"
	"eddy.com/todo/entity"
)

func Insert(todo entity.Todo) error {
	err := data.Insert(todo)
	if err != nil {
		return err
	}

	return nil
}

func Find() ([]entity.Todo, error){
	todoList, err := data.FindAll()
	if err != nil {
		return nil, err
	}

	return todoList, nil
}

func Update(todo entity.Todo) error {
	err := data.Update(todo)
	if err != nil {
		return err
	}

	return nil
}

func Delete(todoId int) error {
	err := data.Delete(todoId)
	if err != nil {
		return err
	}

	return nil
}

func SetToFinished(todoId int) error {
	err := data.SetToFinished(todoId)
	if err != nil {
		return err
	}

	return nil
}
