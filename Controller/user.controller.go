package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
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

	fmt.Println(user.Email)

	dbName := os.Getenv("DB_NAME")
	fmt.Println(dbName)
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
	userID, ok := r.Context().Value("userID").(string)
	fmt.Println(userID)
	if !ok {
		http.Error(w, "Invalid token data", http.StatusInternalServerError)
		return
	}

	db := dbconnection.Connection()
	if db == nil {
		log.Fatal("Some thing is going Wrong")
		return
	}
	connection := dbconnection.Connection()
	if connection == nil {
		http.Error(w, "Some issue in Db Connection", http.StatusInternalServerError)
		return
	}

	var user bson.M
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "I am Unable to get a body ", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	// var response interface{}
	dbName := os.Getenv("DB_NAME")
	collection := db.Database(dbName).Collection("User")
	update := bson.M{"$set": user}
	_, err = collection.UpdateOne(context.Background(), bson.M{
		"email": userID,
	}, update)
	if err != nil {
		http.Error(w, "Error updating data in database", http.StatusInternalServerError)
		return
	}

	// Return a success response
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Data updated successfully"))

}

func UpdateProfilePicture(w http.ResponseWriter, r *http.Request) {
	db := dbconnection.Connection()
	if db == nil {
		log.Fatal("Some thing is going Wrong")
		return
	}

	r.ParseMultipartForm(32 << 20)
	//ParseMultipartForm parses a request body as multipart/form-data
	file, handler, err := r.FormFile("file") //retrieve the file from form data
	//replace file with the key your sent your image with
	if err != nil {
		http.Error(w, "I am Unable to get a body ", http.StatusInternalServerError)
		return
	}
	defer file.Close() //close the file when we finish
	//this is path which  we want to store the file
	f, err := os.OpenFile("/uploads"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		http.Error(w, "I am Unable to get a body ", http.StatusInternalServerError)
		return
	}
	defer f.Close()
	io.Copy(f, file)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(handler.Filename))
}
