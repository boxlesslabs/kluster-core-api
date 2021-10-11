package tests

import (
	"github.com/klusters-core/api/config/db"
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

	requestJSON = `{"phone": "2348131658197", "password": "password"}`

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
	authModel, err = authRepo.GetByPhone(&request.Phone)
	if assert.NoError(t, err) {
		assert.NotEmpty(t, authModel)
	}

	assert.NotEmpty(t, authModel.ID)

	authModel, err = authRepo.ComparePasswords(&authModel.ID, newPassword)
	if assert.NotEmpty(t, authModel) {
		assert.NoError(t, err)
		assert.Equal(t, authModel.Password, newPassword)
	}
}

func TestUpdatePassword(t *testing.T) {
	newPassword := Util.HashPassword("passwords")
	authModel, err = authRepo.GetByPhone(&request.Phone)
	if assert.NoError(t, err) {
		assert.NotEmpty(t, authModel)
	}

	authModel, err = authRepo.UpdatePassword(&authModel.ID, newPassword)
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
		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.NotEmpty(t, rec.Body.String())
		log.Println(rec.Body.String())
	}
}