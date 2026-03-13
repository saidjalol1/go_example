package main

import (
	"fmt"
	"net/http"
	"time"

	"hello/handlers"
	"hello/middlewares"
)


func main() {
	mux := http.NewServeMux()	
	port := "8080"

	fmt.Println("🚀 Server starting...")
	fmt.Println("📍 URL: http://localhost:" + port)
	fmt.Println("🕒 Started at:", time.Now().Format(time.RFC1123))



	mux.HandleFunc("/", handlers.HomeHandler)

	err := http.ListenAndServe(":"+port , middlewares.LoggingMiddleware(mux))

	if err!= nil {
		fmt.Println("❌ Server failed to start:", err)
	}
}
