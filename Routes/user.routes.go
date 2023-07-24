package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	Name  string `json: "name"`
	Email string `json: "email"`
}

func Login(w http.ResponseWriter, r *http.Request) {
	var LoginData = []User{
		{Name: "Almod Milk", Email: "awaisniaz1995@gmail.comss"},
	}
	fmt.Println("Endpoint hit: returnAllGroceries")

	json.NewEncoder(w).Encode(LoginData)
}

func Register(w http.ResponseWriter, r *http.Request) {
	var LoginData = []User{
		{Name: "Almod Milk", Email: "awaisniaz1995@gmail.comss"},
	}
	fmt.Println("Endpoint hit: returnAllGroceries")

	json.NewEncoder(w).Encode(LoginData)
}
