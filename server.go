package main

import (
    "fmt"
    "net/http"
    "github.com/gorilla/mux"
	"github.com/awaisniaz/todo/controller"
)

func main() {
    router := mux.NewRouter()
    router.HandleFunc("/login", controller.Login).Methods("POST")
    router.HandleFunc("/register", controller.Register).Methods("GET")

    // Start the HTTP server
    port := ":3000"
    fmt.Printf("Server is running on port %s...\n", port)
    http.ListenAndServe(port, router)
}

