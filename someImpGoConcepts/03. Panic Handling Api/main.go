package main

import (
	"encoding/json"
	"log"
	"net/http"
	"runtime/debug"
	"time"
)

// Api middlewares should catch panics
// Keeps the server running
// Protects users from seeing stack traces
// Provides consistent error responses.
// Centralized Panic handling

// Using Panic inside handlers is rare but valid when :
// The state is impossible/corrupted.
// You deliberately want to crash only the request, not the server.
// You want to avoid deep error-return chains.

//  MIDDLEWARE PANIC RECOVERY

func RecoverMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("panic recovered: %v \n stack trace : \n%s", err, debug.Stack())

				w.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(w).Encode(map[string]any{
					"error":   "internal server error",
					"message": err,
				})
			}
		}()
		next.ServeHTTP(w, r)
	})
}

// HANDLERS

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello! Everything is fine. \n"))
}

func PanicHandler(w http.ResponseWriter, r *http.Request) {
	// intentionally trigger a panic
	panic("something terrible happened!")
}

func RandomLogicHandler(w http.ResponseWriter, r *http.Request) {
	// Example : Panic on forbidden condition

	x := time.Now().Unix() % 2

	if x == 0 {
		panic("random condition failed")
	}

	w.Write([]byte("Random logic passed. \n"))
}

func main() {
	mux := http.NewServeMux()

	mux.Handle("/hello", RecoverMiddleware(http.HandlerFunc(HelloHandler)))
	mux.Handle("/panic", RecoverMiddleware(http.HandlerFunc(PanicHandler)))
	mux.Handle("/random", RecoverMiddleware(http.HandlerFunc(RandomLogicHandler)))

	log.Println("Server listening on http://localhost:8000")
	log.Fatal(http.ListenAndServe(":8000", mux))
}
