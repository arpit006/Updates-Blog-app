package router

import (
	"github.com/gorilla/mux"
	"log"
)

var ROUTER *mux.Router

func GetHttpRouterFactory() *mux.Router {
	if ROUTER != nil {
		return ROUTER
	}
	log.Println("Router object instantiated on startup!.")
	ROUTER := mux.NewRouter()
	return ROUTER
}