package main

import (
    "fmt"
    "net/http"
    "github.com/gorilla/mux"
	"github.com/awaisniaz/todo/controller"
    "github.com/awaisniaz/todo/dbconnection"

)

func main() {

    router := mux.NewRouter()
    router.HandleFunc("/login", controller.Login).Methods("POST")
    router.HandleFunc("/register", controller.Register).Methods("GET")

    // Start the HTTP server
    port := ":3000"
    fmt.Printf("Server is running on port %s...\n", port)
    http.ListenAndServe(port, router)
    client, ctx, cancel, err := dbconnection.Connect("mongodb://localhost:27017")
    if err != nil{
        panic(err)
    }
     
    // Release resource when the main
    // function is returned.
    defer dbconnection.Close(client, ctx, cancel)
     
    // Ping mongoDB with Ping method
    dbconnection.Ping(client, ctx)
}

