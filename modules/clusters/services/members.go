package services

import (
	authRepo "github.com/klusters-core/api/modules/auth/repo"
	"github.com/klusters-core/api/modules/auth/services"
	"github.com/klusters-core/api/modules/clusters/repo"
	"github.com/labstack/echo"
)

func NewMembersService(service repo.ClusterRepo, auth authRepo.AuthRepo) *membersService {
	return &membersService{ ClusterRepo: service, Auth: services.NewAuthService(auth) }
}

type (
	membersService struct {
		repo.ClusterRepo
		Auth services.UserService
	}

	MembersService interface {
		InviteMembers(ctx echo.Context) error
		RemoveMember(ctx echo.Context) error
		AcceptInvite(ctx echo.Context) error
		GenerateSharingLink(ctx echo.Context) error
		DeactivateLink(ctx echo.Context) error
	}
)

func (m membersService) InviteMembers(ctx echo.Context) error {
	panic("implement me")
}

func (m membersService) RemoveMember(ctx echo.Context) error {
	panic("implement me")
}

func (m membersService) AcceptInvite(ctx echo.Context) error {
	panic("implement me")
}

func (m membersService) GenerateSharingLink(ctx echo.Context) error {
	panic("implement me")
}

func (m membersService) DeactivateLink(ctx echo.Context) error {
	panic("implement me")
}