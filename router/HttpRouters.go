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
	router.HandleFunc("/", services.IndexGetHandler).Methods(http.MethodGet)
	router.HandleFunc("/", services.IndexPostHandler).Methods(http.MethodPost)
	fs := http.FileServer(http.Dir("./static/"))
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))
}

// LoginRouter router for Login Methods
type LoginRouter struct {
}

func (lR LoginRouter) RegisterRoute(router *mux.Router) {
	router.HandleFunc("/login", services.LoginGetHandler).Methods(http.MethodGet)
	router.HandleFunc("/login", services.LoginPostHandler).Methods(http.MethodPost)
	router.HandleFunc("/test", services.LoginTestHandler).Methods(http.MethodGet)
}

// RegisterRouter registers in new user
type RegisterRouter struct {
}

func (rR RegisterRouter) RegisterRoute(router *mux.Router) {
	router.HandleFunc("/register", services.RegisterGetHandler).Methods(http.MethodGet)
	router.HandleFunc("/register", services.RegisterPostHandler).Methods(http.MethodPost)
}

type LogoutRouter struct {
}

func (lR LogoutRouter) RegisterRoute(router *mux.Router) {
	router.HandleFunc("/logout", services.Logout).Methods(http.MethodGet)
}


