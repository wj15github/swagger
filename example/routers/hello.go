package routers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
	"github.com/wj15github/swagger/example/controllers"
)

func SetHelloRoutes(subRouter *mux.Router) {
	hello1APIHandler := negroni.Wrap(http.HandlerFunc(controllers.Hello1))
	subRouter.Handle("/hello1", common.With(hello1APIHandler)).Methods(http.MethodGet)

	hello2APIHandler := negroni.Wrap(http.HandlerFunc(controllers.Hello2))
	subRouter.Handle("/hello2", common.With(hello2APIHandler)).Methods(http.MethodGet)
}
