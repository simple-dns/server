package route

import (
	"github.com/leganck/simple-dns/server/web/handler"
	"net/http"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
	Auth        bool
}

type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		handler.Index,
		true,
	}, Route{
		"Health",
		"GET",
		"/healthz",
		handler.Health,
		false,
	},
	Route{
		"Rpc",
		"POST",
		"/rpc/record",
		handler.RpcAddRecord,
		false,
	},
}
