package controller

import (
	"github.com/gin-gonic/gin"
	"log"
	"openshift-rollouter/config"
	"openshift-rollouter/service"
	"openshift-rollouter/service/openshift"
)

func Init() error {
	port := config.NewConfig().Viper.GetString("server.port")
	if len(port) <= 0 {
		log.Println("set default port")
		port = "8080"
	}
	r := gin.New()
	r.GET("/ping", func(c *gin.Context) {
		service.Ping(c)
	})
	r.POST("/apply/:namespace", func(c *gin.Context) {
		openshift.Apply(c)
	})
	r.GET("/rollout/:namespace/:kind/:name", func(c *gin.Context) {
		openshift.Rollout(c)
	})
	return r.Run(":" + port)
}
