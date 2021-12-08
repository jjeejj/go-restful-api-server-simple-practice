package main

import (
	"errors"
	"go-restful-api-server-simple-practice/router"
	"log"
	"net/http"
	"time"

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
	// ping the server to make ture the server is working
	go func() {
		if err := pingServer(); err != nil {
			log.Fatal("the server router /sd/health has no response, or ir might took too lang to start up", err)
		}
		log.Println("The server has been deployed successfully")
	}()

	log.Printf("Start to listening the incoming requests on http address: %s", ":8080")
	log.Println(http.ListenAndServe(":8080", g).Error())
}

// pingServer ping the server to make ture the server is working
// retry 10 times
func pingServer() error {
	for i := 0; i < 10; i++ {
		resp, err := http.Get("http://127.0.0.1:8080" + "/sd/health")
		if err == nil && resp.StatusCode == 200 {
			return nil
		}
		// Sleep Second to contnue the next ping
		log.Println("Waiting for the server, retry in 1 second")
		time.Sleep(time.Second)
	}
	return errors.New("Cannot connect to the server by /sd/health")
}
