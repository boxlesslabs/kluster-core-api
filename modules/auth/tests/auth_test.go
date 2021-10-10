package tests

import (
	"github.com/klusters-core/api/config/db"
	"github.com/klusters-core/api/modules/auth/models"
	"github.com/klusters-core/api/modules/auth/repo"
	"github.com/klusters-core/api/utils"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

// test variables
var (
	request = &models.AuthModel{
		Phone:     "",
		Password:  "password",
	}

	authModel *models.AuthModel
	err error
	MongoClient = db.Connect()
	authRepo = repo.NewAuthRepo(MongoClient)
	Util utils.GeneralUtil
)

// REPOSITORY TESTS

func TestCreateAndGetAuth(t *testing.T) {
	authModel, err = authRepo.Create(request)
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