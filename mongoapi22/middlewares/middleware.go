package middleware

import (
	"fmt"
	"net/http"
	"time"
)

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()

		fmt.Println("Received request:", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)

		duration := time.Since(startTime)

		fmt.Printf("Request processed in %d ms\n", duration.Milliseconds())
	})
}
