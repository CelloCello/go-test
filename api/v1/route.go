package v1

import (
	"go-test/pkg/net"
	"net/http"
)

func SetRoutes(basePath string) {
	net.RegRoutes(basePath, routeV1)
}

// @title Gin swagger
// @version 1.0
// @description Gin swagger
// @host localhost:8080
// @BasePath /api/v1
// schemes http
var routeV1 = []net.Route{
	{
		Method: http.MethodGet,
		Path:   "ping",
		Handlers: []net.HandlerFunc{
			Ping,
		},
	},
}
