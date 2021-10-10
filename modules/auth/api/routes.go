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

func IndexAuth(path string, router *echo.Group, con db.StartMongoClient) {
	Auth := RegisterAuthService(con)

	// post requests
	router.POST(path, Auth.Authenticate)
	router.POST(path+"forgot-password", Auth.ForgotPassword)
	router.POST(path+"change-password", Auth.ChangePassword, middlewares.IsValidUser(con))

	// get requests
	router.GET(path+"/refresh-token", Auth.RefreshToken, middlewares.IsValidUser(con))
}