package main

import (
	apiv1 "go-test/api/v1"
	"go-test/pkg/net"
	"net/http"
)

var routeV1 = []net.Route{
	{
		Method: http.MethodGet,
		Path:   "ping",
		Handlers: []net.HandlerFunc{
			apiv1.TestPing,
		},
	},
}
