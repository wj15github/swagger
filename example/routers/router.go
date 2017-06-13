package routers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
	"github.com/wj15github/swagger/example/controllers"
)

var (
	common = negroni.New(
		negroni.HandlerFunc(controllers.LoggerMiddleware),
	)
)

func InitRoutes() *mux.Router {
	router := mux.NewRouter()
	router.PathPrefix("/api/").Handler(http.StripPrefix("/api/", http.FileServer(http.Dir("./swagger-ui-2.2.10"))))
	apiRouter := router.PathPrefix("/api/1.0").Subrouter()
	SetHelloRoutes(apiRouter)
	return router
}
