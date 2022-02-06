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
	// logging config override
	SetLoggingConfig()

	http.ListenAndServe(":8080", r)
}

func SetLoggingConfig() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

}

func registerHttpRouters(r *mux.Router) {
	// registering Random Router
	log.Println("Registering Random Router......")
	var randomRouter router.HttpRouter = &router.RandomRouter{}
	randomRouter.RegisterRoute(r)

	//registering Login Router
	log.Println("Registering Login Router......")
	var loginRouter router.HttpRouter = &router.LoginRouter{}
	loginRouter.RegisterRoute(r)

	//registering signup Router
	log.Println("Registering Signup Router......")
	var registerRouter router.HttpRouter = &router.RegisterRouter{}
	registerRouter.RegisterRoute(r)

	// registering Logout Router
	log.Println("Registering Logout Router......")
	var logoutRouter router.HttpRouter = &router.LogoutRouter{}
	logoutRouter.RegisterRoute(r)
}
