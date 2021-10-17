package tests

import (
	"github.com/klusters-core/api/config/db"
	"github.com/klusters-core/api/middlewares"
	"github.com/klusters-core/api/modules/account/repo"
	"github.com/klusters-core/api/modules/account/services"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

var (
	requestJSON = `{"phone": "2348131658199", "password": "password"}`
	MongoClient = db.Connect()
	accountRepo = repo.NewAccountRepo(MongoClient)
	accountService = services.NewAccountService(accountRepo)
	token = `eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRoX2lkIjoiNjE2YjI4Y2E0N2IyMGQ5MDJhY2NhOWYxIiwidXNlcl9pZCI6IjYxNmIyOGNhNDdiMjBkOTAyYWNjYTlmMiIsInBob25lIjoiMjM0ODEzMTY1ODE5OSIsImV4cCI6MTYzNDY3MjA2MH0.S0qq27uAs8ISEFONPrx1vwn7itZ2FWTeuOSK5GZiiX0`
)


// SERVICE TESTS
func TestCreateAccount(t *testing.T) {
	ctx := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(requestJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	req.Header.Set("Content-Type", "application/json")

	log.Println(requestJSON, "request object")

	rec := httptest.NewRecorder()
	httpCtx := ctx.NewContext(req, rec)
	httpCtx.SetPath("/account")

	// assertions
	if assert.NoError(t, accountService.CreateUser(httpCtx)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.NotEmpty(t, rec.Body.String())
		log.Println(rec.Body.String())
	}
}

func TestGetAccount(t *testing.T) {
	ctx := initEcho()
	req := httptest.NewRequest(http.MethodGet, "/account", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()
	ctx.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.NotEmpty(t, rec.Body.String())

	log.Println(rec.Body.String(), "user account")
}


/// HELPERS
func initEcho() *echo.Echo {
	e := echo.New()
	e.Use(middlewares.IsValidUser(MongoClient))
	e.GET("/account", accountService.GetUser)
	return e
}