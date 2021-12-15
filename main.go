package main

import (
	"errors"
	"go-restful-api-server-simple-practice/router"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"

	"go-restful-api-server-simple-practice/config"
)

var (
	// 传入配置文件下地址
	cfg = pflag.StringP("config", "c", "", "apiserver config file path")
)

func main() {
	// init config
	pflag.Parse()
	if err := config.Init(*cfg); err != nil {
		panic(err)
	}
	gin.SetMode(viper.GetString("runmode"))
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
	addr := viper.GetString("addr")
	log.Printf("Start to listening the incoming requests on http address: %s", addr)
	log.Println(http.ListenAndServe(addr, g).Error())
}

// pingServer ping the server to make ture the server is working
// retry 10 times
func pingServer() error {
	for i := 0; i < viper.GetInt("max_ping_count"); i++ {
		resp, err := http.Get(viper.GetString("ping_url") + "/sd/health")
		if err == nil && resp.StatusCode == 200 {
			return nil
		}
		// Sleep Second to contnue the next ping
		log.Println("Waiting for the server, retry in 1 second")
		time.Sleep(time.Second)
	}
	return errors.New("Cannot connect to the server by /sd/health")
}
