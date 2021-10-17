package api

import (
	"errors"
	"github.com/klusters-core/api/integrations/eyowo/repo"
	"github.com/klusters-core/api/integrations/eyowo/services"
	"github.com/robfig/cron/v3"
	"log"
	"github.com/klusters-core/api/integrations/goCache"
)

func RegisterAuthCron(baseUrl string, key string, refreshToken string, appCache goCache.AppCache) {
	eyowoRepo := repo.NewDeveloperRepo(baseUrl, key)
	eyowoService := services.NewAuthService(eyowoRepo, refreshToken, appCache)

	eyowoService.RefreshToken()

	c := cron.New()
	_, err := c.AddFunc("@daily", eyowoService.RefreshToken)
	if err != nil {
		log.Println(err)
		log.Println(errors.New("error encountered starting eyowo auth cron"))
	}
	c.Start()
	log.Println(c.Entries())
	log.Println("cron started")
}