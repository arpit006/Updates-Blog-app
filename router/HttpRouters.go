package router

import (
	"arpit006/web_app_with_go/handlers"
	"arpit006/web_app_with_go/services"
	"github.com/gorilla/mux"
	"net/http"
)

type HttpRouter interface {
	RegisterRoute(router *mux.Router)
}

// UpdatesRouter a random Router
type UpdatesRouter struct {
}

func (rR UpdatesRouter) RegisterRoute(router *mux.Router) {
	router.HandleFunc("/", handlers.Logger(handlers.Authenticator(services.UpdatesGetHandler))).Methods(http.MethodGet)
	router.HandleFunc("/", handlers.Logger(handlers.Authenticator(services.UpdatesPostHandler))).Methods(http.MethodPost)
	fs := http.FileServer(http.Dir("./static/"))
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))
}

// LoginRouter router for Login Methods
type LoginRouter struct {
}

func (lR LoginRouter) RegisterRoute(router *mux.Router) {
	router.HandleFunc("/login", handlers.Logger(services.LoginGetHandler)).Methods(http.MethodGet)
	router.HandleFunc("/login", handlers.Logger(services.LoginPostHandler)).Methods(http.MethodPost)
	router.HandleFunc("/test", handlers.Logger(services.LoginTestHandler)).Methods(http.MethodGet)
}

// RegisterRouter registers in new user
type RegisterRouter struct {
}

func (rR RegisterRouter) RegisterRoute(router *mux.Router) {
	router.HandleFunc("/register", handlers.Logger(services.RegisterGetHandler)).Methods(http.MethodGet)
	router.HandleFunc("/register", handlers.Logger(services.RegisterPostHandler)).Methods(http.MethodPost)
}

// LogoutRouter registers logout activity
type LogoutRouter struct {
}

func (lR LogoutRouter) RegisterRoute(router *mux.Router) {
	router.HandleFunc("/logout", handlers.Logger(services.Logout)).Methods(http.MethodGet)
}

type AuthErrorRouter struct {
}

func (eR AuthErrorRouter) RegisterRoute(router *mux.Router) {
	router.HandleFunc("/auth-error", handlers.Logger(services.AuthErrorHandler)).Methods(http.MethodGet)
}

