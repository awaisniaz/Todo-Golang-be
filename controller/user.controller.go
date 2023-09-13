package controller

import (
	"fmt"
	"net/http"
	"encoding/json"
	"io/ioutil"
	"log"
)

func Login(w http.ResponseWriter, r *http.Request){
	body,err := ioutil.ReadAll(r.Body)
	if err != nil {
        http.Error(w, "Failed to read request body", http.StatusBadRequest)
        return
    }

    // Close the request body to prevent resource leaks
    defer r.Body.Close()

	var user map[string]interface{}

	if err := json.Unmarshal(body, &user); err != nil {
        http.Error(w, "Failed to parse JSON data", http.StatusBadRequest)
        return
    }
	username, ok := user["username"].(string)
	fmt.Printf(username)
	fmt.Print(ok)
	resp := make(map[string]string)
	resp["message"] = "Status Created"
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}
	w.Write(jsonResp)
	return
}

func Register(w http.ResponseWriter, r *http.Request){
	body,err := ioutil.ReadAll(r.Body)
	if err != nil {
        http.Error(w, "Failed to read request body", http.StatusBadRequest)
        return
    }

    // Close the request body to prevent resource leaks
    defer r.Body.Close()

	var user map[string]interface{}

	if err := json.Unmarshal(body, &user); err != nil {
        http.Error(w, "Failed to parse JSON data", http.StatusBadRequest)
        return
    }
	username, ok := user["username"].(string)
	fmt.Printf(username)
	fmt.Print(ok)
	resp := make(map[string]string)
	resp["message"] = "Status Created"
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}
	w.Write(jsonResp)
	return
}