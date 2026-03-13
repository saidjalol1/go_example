package handlers


import (
	"net/http"
	"github.com/gorilla/mux"
)


func BookPageHandler(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	title := vars["title"]
	page := vars["page"]
	
	response := "You've requested the page " + page + " of the book " + title
	w.Write([]byte(response))
}