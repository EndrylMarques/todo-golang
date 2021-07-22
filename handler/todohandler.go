package handler

import (
	"eddy.com/todo/entity"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"eddy.com/todo/service"
)

func Health(responseWriter http.ResponseWriter, request *http.Request) {
	fmt.Fprint(responseWriter, "Server OK")
}

func Insert(responseWriter http.ResponseWriter, request *http.Request) {
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		http.Error(responseWriter, "Error reading requst body", http.StatusInternalServerError)
	}

	var todo entity.Todo
	err = json.Unmarshal(body, &todo)
	if err != nil {
		http.Error(responseWriter, "Error parsing request body", http.StatusInternalServerError)
	}

	err = service.Insert(todo)
	if err != nil {
		http.Error(responseWriter, "Error inserting todo", http.StatusInternalServerError)
	}
}

func GetAll(responseWriter http.ResponseWriter, request *http.Request) {
	todoList, err := service.Find()
	if err != nil {
		http.Error(responseWriter, "Error getting todo list", http.StatusInternalServerError)
	}

	jsonBody, err := json.Marshal(todoList)
	if err != nil {
		http.Error(responseWriter, "Error pasing todo to json", http.StatusInternalServerError)
	}

	responseWriter.Write(jsonBody)
}

func Update(responseWriter http.ResponseWriter, request *http.Request) {
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		http.Error(responseWriter, "Error reading requst body", http.StatusInternalServerError)
	}

	var todo entity.Todo
	err = json.Unmarshal(body, &todo)
	if err != nil {
		http.Error(responseWriter, "Error parsing request body", http.StatusInternalServerError)
	}


	err = service.Update(todo)
	if err != nil {
		http.Error(responseWriter, "Error updating todo", http.StatusInternalServerError)
	}
}

func Delete(responseWriter http.ResponseWriter, request *http.Request) {
	query := request.URL.Query()
	todo := query.Get("todo-id")

	todoID, err := strconv.Atoi(todo)
	if err != nil {
		http.Error(responseWriter, "Error parsing query param", http.StatusInternalServerError)
	}

	err = service.Delete(todoID)
	if err != nil{
		http.Error(responseWriter, "Error deleting todo", http.StatusInternalServerError)
	}
}

func SetFinished(responseWriter http.ResponseWriter, request *http.Request) {
	query := request.URL.Query()
	todo := query.Get("todo-id")

	todoID, err := strconv.Atoi(todo)
	if err != nil {
		http.Error(responseWriter, "Error parsing query param", http.StatusInternalServerError)
	}

	err = service.SetToFinished(todoID)
	if err != nil{
		http.Error(responseWriter, "Error setting todo to finish", http.StatusInternalServerError)
	}
}
