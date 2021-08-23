//=============================================================================
// developer: boxlesslabsng@gmail.com
//=============================================================================

/**
 * Application entry point
 *
 * MODULE - github.com/klusters-core/api
**/

package main

import (
	"fmt"
	"github.com/klusters-core/api/config/db"
	"github.com/klusters-core/api/config/secrets"
	account "github.com/klusters-core/api/modules/account/api"
	auth "github.com/klusters-core/api/modules/auth/api"
	cluster "github.com/klusters-core/api/modules/clusters/api"
	"github.com/klusters-core/api/utils"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/urfave/negroni"
	"log"
	"net/http"
)

// initialize middlewares for access control
func init() {
	response := new(utils.Result)
	middleware.ErrJWTMissing.Code = http.StatusUnauthorized
	middleware.ErrJWTMissing.Message = response.ReturnErrorResult("Could not find authorization token")
}

// initialize a connection to mongo db
var mongoCon = db.Connect()

// RegisterRoutes register routes function to be used by negroni to mount routes
func RegisterRoutes() *echo.Echo {
	router := echo.New()
	router.Use(middleware.Logger())
	router.Use(middleware.Recover())
	router.Use(middleware.CORS())

	rootRouter := router.Group("/klusta/api")

	// initialize routes for account
	account.IndexAccount("/account", rootRouter, mongoCon)
	auth.IndexAuth("/auth", rootRouter, mongoCon)
	cluster.IndexCluster("/cluster", rootRouter, mongoCon)

	return router
}

// main application entry point
func main() {
	fmt.Println("Welcome to kluster api v1.0")

	n := negroni.Classic()
	n.UseHandler(RegisterRoutes())

	err := http.ListenAndServe(":"+secrets.GetSecrets().ApplicationPort, n)
	if err != nil {
		if r := recover(); r != nil {
			err = r.(error)
		}
		log.Fatal(err)
	}
}