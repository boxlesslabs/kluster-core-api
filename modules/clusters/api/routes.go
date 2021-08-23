package api

import (
	"github.com/klusters-core/api/config/db"
	"github.com/labstack/echo"
)

func IndexCluster(path string, router *echo.Group, con db.StartMongoClient) {
	Manager := RegisterManagerService(con)
	Members := RegisterMembersService(con)

	// post requests
	router.POST(path, Manager.AddNewCluster)
	router.POST(path+"/members/:clusterID", Members.InviteMembers)

	// get requests
	router.GET(path+"/:clusterID", Manager.GetSingleCluster)
	router.GET(path+"/members/accept", Members.AcceptInvite)
	router.GET(path+"/members/generate-link/:clusterID", Members.GenerateSharingLink)
	router.GET(path+"/members/deactivate-link/:clusterID", Members.DeactivateLink)

	// delete requests
	router.DELETE(path+"/:clusterID", Manager.DeleteMyCluster)
	router.DELETE(path+"/members/:clusterID", Members.RemoveMember)

	// put requests
	router.PUT(path+"/:clusterID", Manager.EditMyCluster)
	router.PUT(path+"/visibility/:clusterID", Manager.ToggleVisibility)
}
