package main

import (
	"go-restful-api-server-simple-practice/router"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	g := gin.New()

	// middleware
	middlewares := []gin.HandlerFunc{}

	// routes
	router.Load(
		// Cores.
		g,
		middlewares...,
	)
	log.Printf("Start to listening the incoming requests on htt address: %s", ":8080")
	log.Println(http.ListenAndServe(":8080", g).Error())
}
