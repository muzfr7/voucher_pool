package main

import (
	"log"

	"github.com/gin-gonic/gin"
	customerDomain "github.com/muzfr7/voucher_pool/app/domain/customer"

	repository "github.com/muzfr7/voucher_pool/app/infrastructure/repository"
	customerInterface "github.com/muzfr7/voucher_pool/app/interfaces/customer"
	pingInterface "github.com/muzfr7/voucher_pool/app/interfaces/ping"
	customerUsecase "github.com/muzfr7/voucher_pool/app/usecases/customer"
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

	// instantiate repositories
	customerRepo := repository.NewCustomerRepository()

	// instantiate services
	customerSVC := customerUsecase.NewService(customerRepo)

	// instantiate handlers
	pingHandler := pingInterface.NewHandler()
	customerHandler := customerInterface.NewHandler(customerSVC)

	// create routes and run server
	createRoutes(pingHandler, customerHandler)
	runServer()
}

// createRoutes will define app routes.
func createRoutes(pHandler pingInterface.Handler, cHandler customerDomain.Handler) {
	v1 := router.Group(baseRouteV1)
	{
		v1.GET("ping", pHandler.Ping)
		v1.POST("customers", cHandler.Create)
	}
}

// runServer will start the server.
func runServer() {
	// listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	log.Fatalln(router.Run(":8080"))
}
