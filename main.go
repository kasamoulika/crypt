package main

import (
	"flag"
	"fmt"
	"net/http"

	"crypt.com/v2/application/config"
	"crypt.com/v2/application/constants"
	"crypt.com/v2/application/handlers"
	"crypt.com/v2/application/services/external/hitbtc"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

var (
	basePath *string
	env      *string
)

func init() {
	basePath = flag.String(constants.BasePath, constants.DefaultBasePath, "Path to crypto base path")
	env = flag.String(constants.Env, constants.Development, "Application env : prod/dev")
}

func registerPublicRoutes(router *gin.RouterGroup) {
	router.GET("ping", handlers.GetPing)
}

func registerV1Routes(route *gin.RouterGroup) {
	hitbtcClient := hitbtc.NewHitBtcClient()
	currencyHandler := handlers.NewCurrencyHandler(hitbtcClient)
	route.GET("/currency/:symbol", currencyHandler.GetCurrency)
	route.GET("/currency/all", currencyHandler.GetAllCurrency)
}

func main() {
	flag.Parse()
	config.LoadConfig(*basePath, *env)
	SetupRouter()
}

func SetupRouter() {
	router := gin.New()
	rootGroup := router.Group("/")
	v1Group := router.Group("/v1")
	registerPublicRoutes(rootGroup)
	registerV1Routes(v1Group)

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
