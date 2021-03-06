package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/kelseyhightower/envconfig"
	customerDomain "github.com/muzfr7/voucher_pool/app/domain/customer"
	"github.com/muzfr7/voucher_pool/app/infrastructure/db"
	"github.com/muzfr7/voucher_pool/config"

	repository "github.com/muzfr7/voucher_pool/app/infrastructure/repository"
	customerInterface "github.com/muzfr7/voucher_pool/app/interfaces/customer"
	pingInterface "github.com/muzfr7/voucher_pool/app/interfaces/ping"
	customerUsecase "github.com/muzfr7/voucher_pool/app/usecases/customer"
)

const (
	// baseRoute will be used by all routes.
	baseRouteV1 string = "/api/v1/"
)

var (
	// env will contain envs.
	env config.EnvironmentConfig

	// creates a gin router with default middleware: logger and recovery (crash-free) middleware.
	router = gin.Default()
)

func init() {
	// load environment variables from .env file if present
	if _, err := os.Stat("./.env"); err == nil {
		if err = config.Export("./.env"); err != nil {
			log.Fatal(err)
		}
	}

	// load environment variables into EnvironmentConfig struct
	if err := envconfig.Process("", &env); err != nil {
		log.Fatal(err)
	}
}

func main() {
	// instantiate mysql connection
	newDB := db.NewDatabase(env.DBDriver, env.DBUser, env.DBPass, env.DBHost, env.DBPort, env.DBName)
	conn, err := newDB.GetConnection()
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	// instantiate repositories
	customerRepo := repository.NewCustomerRepository(conn)

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
	log.Fatalln(router.Run(fmt.Sprintf(":%s", env.AppPort)))
}
