package route

import (
	"github.com/gorilla/mux"
	"github.com/leganck/simple-dns/server/web/auth"
	"github.com/leganck/simple-dns/server/web/files"
	"github.com/leganck/simple-dns/server/web/log"
	"net/http"
)

func NewRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler

		handler = route.HandlerFunc
		if route.Auth {
			handler = auth.BasicAuth(handler)
		}
		handler = log.Logger(handler)
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}
	router.PathPrefix("/static").Handler(log.Logger(auth.BasicAuth(http.FileServer(http.FS(files.Static)))))
	return router
}
