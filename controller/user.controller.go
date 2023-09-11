package controller

import (
	"fmt"
	"net/http"
)
func Login(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "I am Login Route")
}

func Register(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "I am Register Route")
}