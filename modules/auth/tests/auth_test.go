package tests

import (
	"github.com/klusters-core/api/config/db"
	"github.com/klusters-core/api/middlewares"
	"github.com/klusters-core/api/modules/auth/models"
	"github.com/klusters-core/api/modules/auth/repo"
	"github.com/klusters-core/api/modules/auth/services"
	"github.com/klusters-core/api/utils"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// test variables
var (
	request = models.AuthModel{
		Phone:     "2348131658197",
		Password:  "password",
	}

	requestJSON = `{"phone": "2348131658199", "password": "password"}`
	requestChangePassword = `{"old_password": "password", "new_password": "passwords"}`
	token = `eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRoX2lkIjoiNjE2YjI4Y2E0N2IyMGQ5MDJhY2NhOWYxIiwidXNlcl9pZCI6IjYxNmIyOGNhNDdiMjBkOTAyYWNjYTlmMiIsInBob25lIjoiMjM0ODEzMTY1ODE5OSIsImV4cCI6MTYzNDY3MjA2MH0.S0qq27uAs8ISEFONPrx1vwn7itZ2FWTeuOSK5GZiiX0`

	authModel *models.AuthModel
	err error
	MongoClient = db.Connect()
	authRepo = repo.NewAuthRepo(MongoClient)
	authService = services.NewAuthService(authRepo)
	Util utils.GeneralUtil
)


// REPOSITORY TESTS
func TestCreateAndGetAuth(t *testing.T) {
	authModel, err = authRepo.Create(&request)
	if assert.NoError(t, err) {
		log.Println(authModel, "record created successfully")
		log.Println(err)
	}
}

func TestGetAuthByPhone(t *testing.T) {
	authModel, err = authRepo.GetByPhone(&request.Phone)
	if assert.NoError(t, err) {
		assert.NotEmpty(t, authModel)
		log.Println(authModel, "record retrieved successfully")
	}
}

func TestComparePassword(t *testing.T) {
	newPassword := Util.HashPassword("password")
	oldPassword := Util.HashPassword("passwords")
	authModel, err = authRepo.GetByPhone(&request.Phone)
	if assert.NoError(t, err) {
		assert.NotEmpty(t, authModel)
	}

	assert.NotEmpty(t, authModel.ID)

	authModel, err = authRepo.ComparePasswords(&authModel.ID, oldPassword, newPassword)
	if assert.NotEmpty(t, authModel) {
		assert.NoError(t, err)
	}
}

func TestUpdatePassword(t *testing.T) {
	newPassword := Util.HashPassword("password")
	oldPassword := Util.HashPassword("passwords")
	authModel, err = authRepo.GetByPhone(&request.Phone)
	if assert.NoError(t, err) {
		assert.NotEmpty(t, authModel)
	}

	authModel, err = authRepo.UpdatePassword(&authModel.ID, newPassword, oldPassword)
	if assert.NotEmpty(t, authModel) {
		assert.NoError(t, err)
		log.Println(authModel, "record retrieved successfully")
	}
}


// SERVICE TESTS
func TestAuthenticate(t *testing.T) {
	ctx := echo.New()
	//requestObj, _ := json.Marshal(request)
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(requestJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	req.Header.Set("Content-Type", "application/json")

	log.Println(requestJSON, "request object")

	rec := httptest.NewRecorder()
	httpCtx := ctx.NewContext(req, rec)
	httpCtx.SetPath("/auth")

	// assertions
	if assert.NoError(t, authService.Authenticate(httpCtx)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.NotEmpty(t, rec.Body.String())
		log.Println(rec.Body.String())
	}
}

func TestRefreshToken(t *testing.T) {
	ctx := initEcho()
	req := httptest.NewRequest(http.MethodGet, "/auth/refresh-token", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()
	ctx.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.NotEmpty(t, rec.Body.String())

	log.Println(rec.Body.String(), "user account")
}

func TestForgotPassword(t *testing.T) {
	ctx := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(requestJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	req.Header.Set("Content-Type", "application/json")

	log.Println(requestJSON, "request object")

	rec := httptest.NewRecorder()
	httpCtx := ctx.NewContext(req, rec)

	// assertions
	if assert.NoError(t, authService.ForgotPassword(httpCtx)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.NotEmpty(t, rec.Body.String())
		log.Println(rec.Body.String())
	}
}

func TestChangePassword(t *testing.T) {
	ctx := initEcho()
	req := httptest.NewRequest(http.MethodPost, "/auth/change-password", strings.NewReader(requestChangePassword))
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
	e.GET("auth/refresh-token", authService.RefreshToken)
	e.POST("auth/change-password", authService.ChangePassword)
	return e
}