package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	dbconnection "github.com/awaisniaz/todo/DbConnection"
	"github.com/awaisniaz/todo/utils"
)

type User struct {
	Name     string `json: "name"`
	Email    string `json: "email"`
	Password string `json:"password"`
}

type Response1 struct {
	Message string `json:message`
}
type Response struct {
	Message error `json:message`
}

func Login(w http.ResponseWriter, r *http.Request) {
	var LoginData = []User{
		{Name: "Almod Milk", Email: "awaisniaz1995@gmail.comss"},
	}
	fmt.Println("Endpoint hit: returnAllGroceries")

	json.NewEncoder(w).Encode(LoginData)
}

func Register(w http.ResponseWriter, r *http.Request) {
	db := dbconnection.Connection()
	if db == nil {
		log.Fatal("Some thing is going Wrong")
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()
	var user User

	// Unmarshal the JSON data into the User struct
	err = json.Unmarshal(body, &user)
	if err != nil {
		http.Error(w, "Failed to unmarshal JSON", http.StatusInternalServerError)
		return
	}
	// Do something with the request body (in this example, just printing it)
	var response interface{}
	dbName := os.Getenv("DB_NAME")
	fmt.Println(dbName)
	password, err := utils.HashPassword(user.Password)
	if err != nil {
		http.Error(w, "Some thing went wrong to convert password", http.StatusInternalServerError)
		return
	}
	user.Password = password
	collection := db.Database(dbName).Collection("User")

	data, err := collection.InsertOne(context.Background(), user)
	fmt.Println(data)
	if err != nil {
		response = Response{
			Message: err,
		}
	} else {
		response = Response1{
			Message: "you are Register Successfully",
		}
	}

	json.NewEncoder(w).Encode(response)

	fmt.Println("Endpoint hit: returnAllGroceries")

}
