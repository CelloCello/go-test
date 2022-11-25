package main

import (
	docs "go-test/app/apiserver/docs"
	"go-test/pkg/net"
)

// @title Gin swagger
// @version 1.0
// @description Gin swagger
// @contact.name Sero
// @contact.url https://ystock.tw/
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @BasePath /api/v1
// schemes http
func main() {
	g := net.Init()
	basePath := "/api/v1"
	net.RegRoutes(basePath, routeV1)
	net.SetSwagger(basePath, "/swagger/*any", docs.SwaggerInfo)
	err := g.Run()
	if err != nil {
		panic("Oops!")
	}
}
