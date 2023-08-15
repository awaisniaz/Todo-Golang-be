package main

import (
	"log"
	"net/http"
	"os"

	controller "github.com/awaisniaz/todo/Controller"
	dbconnection "github.com/awaisniaz/todo/DbConnection"
	middleware "github.com/awaisniaz/todo/Middleware"

	"github.com/gorilla/mux"
)

func main() {
	os.Setenv("DB_NAME", "todo-golang")
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
	r.HandleFunc("/login", controller.Login).Methods("POST")
	r.HandleFunc("/signup", controller.Register).Methods("POST")
	r.Handle("/updateProfile", middleware.Authenticate(http.HandlerFunc(controller.UpdateProfile))).Methods("PATCH")
	r.Handle("/updateProfilePic", middleware.Authenticate(http.HandlerFunc(controller.UpdateProfilePicture))).Methods("PATCH")
	r.HandleFunc("/addTodo", controller.AddTodo)
	r.HandleFunc("/getTodos", controller.GetTodo)
	r.HandleFunc("/updateTodo", controller.UpdateTodo)
	r.HandleFunc("deleteTodo", controller.DeleteTodo)
	log.Fatal(http.ListenAndServe(":10000", r))
}
