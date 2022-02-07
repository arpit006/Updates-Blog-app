package handlers

import (
	"log"
	"net/http"
	"time"
)

// HttpLogger this handler logs all the http request
// https://drstearns.github.io/tutorials/gomiddleware/
type HttpLogger struct {
	handler http.Handler
}

func (lH *HttpLogger) ServeHttp(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	lH.handler.ServeHTTP(w, r)
	log.Printf("WebAppWithGo :: %s %s %v", r.Method, r.URL.Path, time.Since(start))
}

func NewHttpLogger(handler http.Handler) *HttpLogger {
	return &HttpLogger{
		handler: handler,
	}
}
