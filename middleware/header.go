// header 设置响应头
package middleware

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// NoCache is a middleware function that append headers
// to prevent the client from caching the http response
func NoCache(c *gin.Context) {
	c.Header("Cache-Control", "no-cache")
	c.Header("Expired", time.Now().UTC().Format(http.TimeFormat))
	c.Header("Last-Modified", time.Now().UTC().Format(http.TimeFormat))
	c.Next()
}

// Options is a middleware function option request
func Options(c *gin.Context) {
	if c.Request.Method == "OPTIONS" {
		c.Header("Access-Control-Allow-Origin", c.Request.Host)
		c.Header("Access-Control-Allow-Method", "GET,POST,PUT,DELETE")
		c.Header("Access-Control-Allow-Headers", "authorization, orign, content-type, accept")
		c.Header("Content-Type", "application/json")
		c.AbortWithStatus(200)
	}
	c.Next()
}

func Secure(c *gin.Context) {
	c.Next()
}
