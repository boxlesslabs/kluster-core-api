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
	"github.com/klusters-core/api/modules/auth/repo"
	"github.com/klusters-core/api/modules/auth/services"
)

func RegisterAuthService(con db.StartMongoClient) services.UserService {
	authRepo := repo.NewAuthRepo(con)
	return services.NewAuthService(authRepo)
}