package route

import (
	"net/http"

	"eddy.com/todo/handler"
)

func RegisterRoute() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler.Health)
	mux.HandleFunc("/todo/get", handler.GetTodo)
	mux.HandleFunc("/todo/add", handler.InsertTodo)
	mux.HandleFunc("/todo/set-finished", handler.SetoTodoFinishd)

	return mux
}
