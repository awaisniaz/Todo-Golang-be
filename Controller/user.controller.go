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
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	Name     string `json: "name"`
	Email    string `json: "email"`
	Password string `json:"password"`
	Profile  string `json:"profile"`
	Token    string `json:"token"`
}

type Response1 struct {
	Message string `json:message`
}
type Response struct {
	Message error `json:message`
}

func Login(w http.ResponseWriter, r *http.Request) {
	db := dbconnection.Connection()
	if db == nil {
		log.Fatal("Some thing is going Wrong")
		return
	}

	// body, err := ioutil.ReadAll(r.Body)
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	// Unmarshal the JSON data into the User struct
	// err = json.Unmarshal(body, &user)
	// if err != nil {
	// 	http.Error(w, "Failed to unmarshal JSON", http.StatusInternalServerError)
	// 	return
	// }
	fmt.Println(user.Email)

	dbName := os.Getenv("DB_NAME")
	fmt.Println(dbName)
	// filter := bson.D{{"email": user.Email}}
	var findUser User
	collection := db.Database(dbName).Collection("User")
	err = collection.FindOne(context.Background(), bson.M{"email": user.Email}).Decode(&findUser)
	if err != nil {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	passwordVerfication := utils.CheckPasswordHash(user.Password, findUser.Password)

	if passwordVerfication == false {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	token, err := utils.GenerateToken(findUser.Email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	findUser.Token = token
	var LoginData = findUser
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

func UpdateProfile(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w)

}
