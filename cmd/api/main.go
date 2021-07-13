package main

import (
	"log"

	"github.com/gin-gonic/gin"
	pingInterface "github.com/muzfr7/voucher_pool/app/interfaces/ping"
)

var (
	// baseRoute will be used by all routes.
	baseRouteV1 string

	// creates a gin router with default middleware: logger and recovery (crash-free) middleware.
	router = gin.Default()
)

func init() {
	baseRouteV1 = "/api/v1/"
}

func main() {

	// instantiate handlers
	pingHandler := pingInterface.NewHandler()

	// create routes and run server
	createRoutes(pingHandler)
	runServer()
}

// createRoutes will define app routes.
func createRoutes(pHandler pingInterface.Handler) {
	v1 := router.Group(baseRouteV1)
	{
		v1.GET("ping", pHandler.Ping)
	}
}

// runServer will start the server.
func runServer() {
	// listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	log.Fatalln(router.Run(":8080"))
}
