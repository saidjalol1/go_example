package handlers

import "net/http"


// Handler Function for the Home Page 
func HomeHandler(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Welcome to the Home Page!"))
}

