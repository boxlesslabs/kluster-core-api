//=============================================================================
// developer: boxlesslabsng@gmail.com
//=============================================================================

/**
 * Define module routes
 * path sets the parent string for concatenation
 * router points to an echo group for typically grouping related routes
**/

package api

import (
	"github.com/klusters-core/api/config/db"
	"github.com/klusters-core/api/middlewares"
	"github.com/labstack/echo"
)

func IndexAccount(path string, router *echo.Group, con db.StartMongoClient) {
	// http post methods
	Account := RegisterAccountService(con)
	router.POST(path, Account.CreateUser)

	// http get methods
	router.GET(path, Account.GetUser, middlewares.IsValidUser(con))
}