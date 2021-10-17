package tests

import (
	"encoding/json"
	"github.com/klusters-core/api/config/db"
	"github.com/klusters-core/api/middlewares"
	model "github.com/klusters-core/api/modules/account/models"
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
	requestProfile = `{"full_name": "Patrick Jesam", "image_url": "https://passid-media-prod.s3.amazonaws.com/2NP6wxmFY1bb", "email": "patrick@softcom.xyz"}`
	MongoClient = db.Connect()
	accountRepo = repo.NewAccountRepo(MongoClient)
	accountService = services.NewAccountService(accountRepo)
	token = `eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRoX2lkIjoiNjE2YjI4Y2E0N2IyMGQ5MDJhY2NhOWYxIiwidXNlcl9pZCI6IjYxNmIyOGNhNDdiMjBkOTAyYWNjYTlmMiIsInBob25lIjoiMjM0ODEzMTY1ODE5OSIsImV4cCI6MTYzNDY3MjA2MH0.S0qq27uAs8ISEFONPrx1vwn7itZ2FWTeuOSK5GZiiX0`
)

type (
	APIResponse struct {
		Success 		bool				`json:"success"`
		Data 			model.AccountsModel	`json:"data"`
		Message			string				`json:"message"`
	}
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

func TestUpdateProfile(t *testing.T) {
	ctx := initEcho()
	req := httptest.NewRequest(http.MethodPut, "/account", strings.NewReader(requestProfile))
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")

	log.Println(requestProfile, "request object")
	rec := httptest.NewRecorder()
	ctx.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.NotEmpty(t, rec.Body.String())

	log.Println(rec.Body.String(), "user account")
}

func TestToggleAccountStatus(t *testing.T) {
	ctx := initEcho()
	req := httptest.NewRequest(http.MethodPut, "/account/deactivate", strings.NewReader(`{"status": "activated"}`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()
	ctx.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.NotEmpty(t, rec.Body.String())

	log.Println(rec.Body.String())
	var response *APIResponse
	err := json.NewDecoder(rec.Body).Decode(&response)
	log.Println(response, err)
	assert.Equal(t, model.StatusActive, response.Data.Status)
}

func TestUpdateFCMID(t *testing.T) {
	ctx := initEcho()
	req := httptest.NewRequest(http.MethodPut, "/account/fcm", strings.NewReader(`{"fcm_id": "jkf_09kdj67838$%bhha09KJAbamd_ABC"}`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()
	ctx.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.NotEmpty(t, rec.Body.String())

	log.Println(rec.Body.String())
	var response *APIResponse
	err := json.NewDecoder(rec.Body).Decode(&response)
	log.Println(response, err)
}

/// HELPERS
func initEcho() *echo.Echo {
	e := echo.New()
	e.Use(middlewares.IsValidUser(MongoClient))
	e.GET("/account", accountService.GetUser)
	e.PUT("/account", accountService.UpdateProfile)
	e.PUT("/account/deactivate", accountService.ToggleUserStatus)
	e.PUT("/account/fcm", accountService.UpdateFCMID)
	return e
}