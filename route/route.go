package route

import (
	"github.com/gorilla/mux"
	"net/http"

	"eddy.com/todo/handler"
)

func RegisterRoute() http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/", handler.Health).Methods("GET")
	r.HandleFunc("/todo", handler.GetAll).Methods("GET")
	r.HandleFunc("/todo", handler.Insert).Methods("POST")
	r.HandleFunc("/todo", handler.Update).Methods("PUT")
	r.HandleFunc("/todo", handler.Delete).Methods("DELETE")
	r.HandleFunc("/todo/set-finished", handler.SetFinished).Methods("PUT")

	return r
}
