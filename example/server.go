// @APIVersion 1.0
// @APITitle Example API Doc
// @APIDescription this is an example for the swagger doc
// @BasePath http://localhost:3000/api/
package main

import (
	"net/http"

	"github.com/wj15github/swagger/example/routers"

	"github.com/leesper/holmes"
	"github.com/urfave/negroni"
)

func main() {
	defer holmes.Start().Stop()

	router := routers.InitRoutes()
	n := negroni.New()
	n.UseHandler(router)

	holmes.Errorln(http.ListenAndServe(":3000", n))
}
