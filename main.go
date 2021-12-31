package main

import (
	"errors"
	"go-restful-api-server-simple-practice/model"
	"go-restful-api-server-simple-practice/router"

	"github.com/lexkong/log"

	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"

	"go-restful-api-server-simple-practice/config"

	_ "go-restful-api-server-simple-practice/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var (
	// 传入配置文件下地址
	cfg = pflag.StringP("config", "c", "", "apiserver config file path")
)

// @title RsstAPI 练习
// @version 1.0
// @description 描述
// @host 127.0.0.1:8888
func main() {
	// init config
	pflag.Parse()
	if err := config.Init(*cfg); err != nil {
		panic(err)
	}
	// for {
	// 	fmt.Println(viper.GetString("runmode"))
	// 	time.Sleep(4 * time.Second)
	// }
	gin.SetMode(viper.GetString("runmode"))

	// init db
	model.DB.Init()

	g := gin.New()

	// middleware
	middlewares := []gin.HandlerFunc{}
	// swagger
	g.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
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
		log.Info("The server has been deployed successfully")
	}()
	addr := viper.GetString("addr")
	log.Infof("Start to listening the incoming requests on http address: %s", addr)
	log.Info(http.ListenAndServe(addr, g).Error())
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
		log.Info("Waiting for the server, retry in 1 second")
		time.Sleep(time.Second)
	}
	return errors.New("Cannot connect to the server by /sd/health")
}
