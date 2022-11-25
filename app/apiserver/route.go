package main

import (
	apiv1 "go-test/api/v1"
	"go-test/pkg/net"
	"net/http"
)

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
			apiv1.TestPing,
		},
	},
}
