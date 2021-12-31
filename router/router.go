package router

import (
	"go-restful-api-server-simple-practice/handler/sd"
	"go-restful-api-server-simple-practice/handler/user"
	"go-restful-api-server-simple-practice/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Load(g *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {
	g.Use(gin.Recovery())

	g.Use(middleware.NoCache)
	g.Use(middleware.Options)
	g.Use(middleware.Secure)
	g.Use(mw...)

	// 404 handle
	g.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "The incorrect API route.")
	})

	// sd group
	sdR := g.Group("/sd")
	{
		sdR.GET("/health", sd.HealthCheck)
		sdR.GET("/disk", sd.DiskCheck)
		sdR.GET("/cpu", sd.CPUCheck)
		sdR.GET("/ram", sd.RAMCheck)
	}
	// user group
	u := g.Group("/v1/user")
	{
		u.POST("/:username", user.Create)
	}

	return g
}
