package main

import (
	"log"
	"net/http"

	dbconnection "github.com/awaisniaz/todo/DbConnection"
	routes "github.com/awaisniaz/todo/Routes"
	"github.com/gorilla/mux"
)

func main() {
	client, ctx, cancel, err := dbconnection.Connect("mongodb://localhost:27017/todo-golang")
	if err != nil {
		panic(err)
	}
	// Release resource when the main
	// function is returned.
	defer dbconnection.Close(client, ctx, cancel)

	// Ping mongoDB with Ping method
	dbconnection.Ping(client, ctx)
	r := mux.NewRouter()
	r.HandleFunc("/login", routes.Login).Methods("POST")
	r.HandleFunc("/signup", routes.Register).Methods("POST")
	r.HandleFunc("/addTodo", routes.AddTodo)
	r.HandleFunc("/getTodos", routes.GetTodo)
	r.HandleFunc("/updateTodo", routes.UpdateTodo)
	r.HandleFunc("deleteTodo", routes.DeleteTodo)
	log.Fatal(http.ListenAndServe(":10000", r))
}
