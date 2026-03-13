package main

import (
	"net/http"
	"github.com/gorilla/mux"

	"routing/handlers"
)


func main(){
	
	router := mux.NewRouter()
	router.HandleFunc("/books/{title}/page/{page}", handlers.BookPageHandler).Methods("GET")

	http.ListenAndServe(":8080", router)
}