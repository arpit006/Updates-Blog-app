package handlers

import (
	"arpit006/web_app_with_go/sessions"
	"log"
	"net/http"
	"time"
)

// Middlewares in golang

// Logger this handler Logs all the http requests
func Logger(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		handler.ServeHTTP(w, r)
		log.Printf("WebAppWithGo :: %s %s %v", r.Method, r.URL.Path, time.Since(start))
	}
}

// Authenticator this handler authenticates each http request for authentication
func Authenticator(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		sessions.ValidateSession(w, r)
		handler.ServeHTTP(w, r)
	}
}
