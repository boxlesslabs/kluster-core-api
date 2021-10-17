package services

import (
	"github.com/klusters-core/api/integrations/eyowo/repo"
	"github.com/klusters-core/api/integrations/goCache"
	"log"
	"time"
)

type (
	IAuthService interface {
		Authenticate() error
		RefreshToken()
	}

	authService struct {
		Repo			repo.IDeveloperRepo
		Token			string
		AppCache		goCache.AppCache
	}
)

func NewAuthService(repo repo.IDeveloperRepo, refreshToken string, appCache goCache.AppCache) *authService {
	return &authService{Repo: repo, Token:refreshToken, AppCache:appCache}
}

func (service *authService) Authenticate() (err error) {
	result, err := service.Repo.ValidateApp()
	if err != nil {
		return err
	}

	err = service.AppCache.Set("userAuthData", result.Data, 12*time.Hour)
	if err != nil {
		log.Println(err)
	}

	return nil
}

func (service *authService) RefreshToken() {
	result, err := service.Repo.RefreshToken(service.Token)
	if err != nil {
		log.Println(err)
	}

	log.Println(result)

	err = service.AppCache.Set("accessToken", result.Data.AccessToken, 12*time.Hour)
	if err != nil {
		log.Println(err)
	}
}