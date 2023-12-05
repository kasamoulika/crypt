package main

import (
	"fmt"
	"net/http"

	"crypt.com/v2/application/handlers"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func registerPublicRoutes(router *gin.RouterGroup) {
	router.GET("ping", handlers.PingHandler)
}

func main() {
	SetupRouter()
}

func SetupRouter() {
	router := gin.New()
	rootGroup := router.Group("/")
	registerPublicRoutes(rootGroup)

	addr := fmt.Sprintf(":%d", 8080)
	s := &http.Server{
		Addr:    addr,
		Handler: router,
	}

	log.Infof("crypt started on 0.0.0.0%s", addr)
	if err := s.ListenAndServe(); err != nil {
		panic(err)
	}
}
