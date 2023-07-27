package controller

import "net/http"

func AddTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")

	// Write the string response to the ResponseWriter
	// Here, we're sending the string "Hello, world!"
	w.Write([]byte("Hello, world!"))
}

func UpdateTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")

	// Write the string response to the ResponseWriter
	// Here, we're sending the string "Hello, world!"
	w.Write([]byte("Hello, world!"))
}

func GetTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")

	// Write the string response to the ResponseWriter
	// Here, we're sending the string "Hello, world!"
	w.Write([]byte("Hello, world!"))
}

func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")

	// Write the string response to the ResponseWriter
	// Here, we're sending the string "Hello, world!"
	w.Write([]byte("Hello, world!"))
}
