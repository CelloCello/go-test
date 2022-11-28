package v1

import (
	"go-test/pkg/net"
	"net/http"
)

func SetRoutes(basePath string) {
	net.RegRoutes(basePath, routeV1)
}

var routeV1 = []net.Route{
	{
		Method: http.MethodGet,
		Path:   "ping",
		Handlers: []net.HandlerFunc{
			Ping,
		},
	},
	{
		Method: http.MethodGet,
		Path:   "users",
		Handlers: []net.HandlerFunc{
			GetUsers,
		},
	},
	{
		Method: http.MethodPost,
		Path:   "users",
		Handlers: []net.HandlerFunc{
			CreateUser,
		},
	},
}
