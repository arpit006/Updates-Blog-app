package starter

import (
	"arpit006/web_app_with_go/router"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func StartWebApp() {
	log.Println("Starting Web application in go.......")
	r := router.GetHttpRouterFactory()
	// registering routers
	registerHttpRouters(r)
	http.ListenAndServe(":8080", r)
}

func registerHttpRouters(r *mux.Router) {
	// registering Random Router
	log.Println("Registering Random Router......")
	var randomRouter router.HttpRouter = &router.RandomRouter{}
	randomRouter.RegisterRoute(r)
}
