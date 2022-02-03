package router

import (
	"arpit006/web_app_with_go/services"
	"github.com/gorilla/mux"
	"net/http"
)

type HttpRouter interface {
	RegisterRoute(router *mux.Router)
}

// RandomRouter a random Router
type RandomRouter struct {
}

func (rR RandomRouter) RegisterRoute(router *mux.Router) {
	router.HandleFunc("/hello", services.IndexHandler).Methods(http.MethodGet)
}


