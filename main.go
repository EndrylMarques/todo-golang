package main

import (
	"log"
	"net/http"

	"eddy.com/todo/data"
	"eddy.com/todo/route"
)

func main() {

	data.InitDatabase()

	log.Print("Server on")
	http.ListenAndServe(":9000", route.RegisterRoute())
}
