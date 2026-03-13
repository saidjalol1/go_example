package main

import (
	"net/http"
	"http_server/handlers"
)


func main(){	
	var PORT = "8080"

	server := http.NewServeMux()

	server.HandleFunc("/", handlers.Home)

	fs := http.FileServer(http.Dir("./static"))
	server.Handle("/static/", http.StripPrefix("/static/", fs))
	http.ListenAndServe(":"+PORT, server)
}