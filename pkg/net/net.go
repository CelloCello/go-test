package net

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type HandlerFunc = gin.HandlerFunc

type Route struct {
	Method   string
	Path     string
	Handlers []HandlerFunc
}

var ginEngine *gin.Engine

func Config(g *gin.Engine) *gin.Engine {
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AllowBrowserExtensions = true
	corsConfig.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	// corsConfig.AllowHeaders = []string{"Authorization", "Origin"}
	corsConfig.AllowHeaders = []string{"Authorization", "Content-Type", "Upgrade", "Origin",
		"Connection", "Accept-Encoding", "Accept-Language", "Host", "Access-Control-Request-Method", "Access-Control-Request-Headers"}
	g.Use(cors.New(corsConfig))
	return g
}

func Init() *gin.Engine {
	g := gin.Default()
	Config(g)
	ginEngine = g
	return g
}

func RegRoutes(group string, routes []Route) {
	if ginEngine == nil {
		panic("881")
	}

	g := ginEngine.Group(group)
	for _, r := range routes {
		g.Handle(r.Method, r.Path, r.Handlers...)
	}
}
