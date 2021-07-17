package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"eddy.com/todo/entity"
	"eddy.com/todo/service"
)

func Health(responseWriter http.ResponseWriter, request *http.Request) {
	fmt.Fprint(responseWriter, "Server OK")
}

func GetTodo(responseWriter http.ResponseWriter, request *http.Request) {
	if request.Method == "GET" {

		jsonBody, err := json.Marshal(service.Find())
		if err != nil {
			http.Error(responseWriter, "Error pasing todo to json", http.StatusInternalServerError)
		}

		responseWriter.Write(jsonBody)
	} else {
		http.Error(responseWriter, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

func InsertTodo(responseWriter http.ResponseWriter, request *http.Request) {
	if request.Method == "POST" {
		body, err := ioutil.ReadAll(request.Body)
		if err != nil {
			http.Error(responseWriter, "Error reading requst body", http.StatusInternalServerError)
		}

		var todo entity.Todo
		err = json.Unmarshal(body, &todo)
		if err != nil {
			http.Error(responseWriter, "Error parsing request body", http.StatusInternalServerError)
		}

		service.Insert(todo)

	} else {
		http.Error(responseWriter, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

func SetoTodoFinishd(responseWriter http.ResponseWriter, request *http.Request) {
	if request.Method == "PUT" {
		query := request.URL.Query()
		todo := query.Get("todo-id")

		todoID, err := strconv.Atoi(todo)
		if err != nil {
			http.Error(responseWriter, "Error parsing query param", http.StatusInternalServerError)
		}

		service.SetToFinished(todoID)

	} else {
		http.Error(responseWriter, "Invalid request method", http.StatusMethodNotAllowed)
	}
}
