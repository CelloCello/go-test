package main

import (
	docs "go-test/app/apiserver/docs"
	apiv1 "go-test/internal/api/v1"
	"go-test/internal/migrations"
	"go-test/pkg/database"
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
	// database
	database.Init("data/data.sqlite")
	migrations.Init()

	// net engine
	g := net.Init()
	basePath := "/api/v1"
	apiv1.SetRoutes(basePath)
	net.SetSwagger(basePath, "/swagger/*any", docs.SwaggerInfo)
	err := g.Run()
	if err != nil {
		panic("Oops!")
	}
}
