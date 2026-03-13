package middlewares

import (
	"fmt"
	"time"
	"net/http"
)

// Middleware for logging incoming HTTP requests
func LoggingMiddleware(next http.Handler) http.Handler {
	
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		
		start := time.Now()

		fmt.Printf("Started %s at %v\n", r.URL.Path, start)
		fmt.Println()
		fmt.Printf("Client IP: %s\n", r.RemoteAddr)
		fmt.Println()
		fmt.Printf("➡️  Method: %s, URL: %s\n", r.Method, r.URL.Path)
		fmt.Println()

		next.ServeHTTP(w, r)
		

		duration := time.Since(start)
		ms := float64(duration) / float64(time.Millisecond)
		fmt.Printf("⬅️ Completed in %.3f ms\n", ms)
	})
}

