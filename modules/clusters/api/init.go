package api

import (
	"github.com/klusters-core/api/config/db"
	auth "github.com/klusters-core/api/modules/auth/repo"
	"github.com/klusters-core/api/modules/clusters/repo"
	"github.com/klusters-core/api/modules/clusters/services"
)

func RegisterManagerService(con db.StartMongoClient) services.ManagerService {
	clusterRepo := repo.NewClusterRepo(con)
	authRepo := auth.NewAuthRepo(con)
	return services.NewManagerService(clusterRepo, authRepo)
}

func RegisterMembersService(con db.StartMongoClient) services.MembersService {
	clusterRepo := repo.NewClusterRepo(con)
	authRepo := auth.NewAuthRepo(con)
	return services.NewMembersService(clusterRepo, authRepo)
}