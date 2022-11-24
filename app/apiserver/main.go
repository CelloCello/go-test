package main

import (
	"go-test/pkg/net"
)

func main() {
	// r := gin.Default()
	// r.GET("/ping", func(c *gin.Context) {
	//   c.JSON(http.StatusOK, gin.H{
	// 	"message": "pong",
	//   })
	// })
	// r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	g := net.Init()
	net.RegRoutes("api/v1", routeV1)
	err := g.Run()
	if err != nil {
		panic("Oops!")
	}
}
