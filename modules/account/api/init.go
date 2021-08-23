//=============================================================================
// developer: boxlesslabsng@gmail.com
//=============================================================================

/**
 * Register a service for different use cases
 * function is mounted in routes.go
 * function shares same package with init.go
**/

package api

import (
	"github.com/klusters-core/api/config/db"
	"github.com/klusters-core/api/modules/account/repo"
	"github.com/klusters-core/api/modules/account/services"
)

func RegisterAccountService(con db.StartMongoClient) services.UserInterface {
	accountRepo := repo.NewAccountRepo(con)
	return services.NewUserService(accountRepo)
}