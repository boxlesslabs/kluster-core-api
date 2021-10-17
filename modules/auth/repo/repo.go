package repo

import (
	"errors"
	"github.com/klusters-core/api/config/db"
	"github.com/klusters-core/api/config/error_response"
	"github.com/klusters-core/api/modules/auth/models"
	"github.com/klusters-core/api/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

type (
	authRepo struct {
		col 	db.MongoInterface
		client 	db.StartMongoClient
		authRes models.AuthModel
		errRes error
		utils.GeneralUtil
	}

	AuthRepo interface {
		Create(request *models.AuthModel) (*models.AuthModel, error)
		GetByCredentials(request *models.AuthModel) (*models.AuthModel, error)
		GetByPhone(phone *string) (*models.AuthModel, error)
		ReturnClient() db.StartMongoClient
		UpdatePassword(id *primitive.ObjectID, password string) (*models.AuthModel, error)
		ComparePasswords(userID *primitive.ObjectID, password string) (*models.AuthModel, error)
		GetByID(id *primitive.ObjectID) (*models.AuthModel, error)
	}
)


// CONSTRUCTOR
func NewAuthRepo(client db.StartMongoClient) *authRepo {
	col := db.NewMongoCollection("auth", client)
	return &authRepo{col: col, client:client}
}


// PUBLIC
func (auth *authRepo) Create(request *models.AuthModel) (*models.AuthModel, error) {
	if _, auth.errRes = auth.GetByPhone(&request.Phone); auth.errRes == nil {
		return nil, errors.New(error_response.DuplicateError{Resource:"auth"}.Error())
	}

	request.Password = auth.HashPassword(request.Password)
	request.NewID()
	request.CreatedStamp()
	if auth.authRes.ID, auth.errRes = auth.col.AddSingleReturnID(request); auth.errRes != nil {
		return nil, auth.errRes
	}
	return auth.GetByID(&auth.authRes.ID)
}

func (auth *authRepo) UpdatePassword(id *primitive.ObjectID, password string) (*models.AuthModel, error) {
	log.Println(id, password, "request gotten from service")
	result, err := auth.col.UpdateById(*id, bson.D{{"password", auth.HashPassword(password)}})
	if err != nil {
		return nil, err
	}

	return auth.DecodeSingle(result)
}

func (auth *authRepo) ComparePasswords(userID *primitive.ObjectID, password string) (*models.AuthModel, error) {
	return auth.DecodeSingle(auth.col.GetSingleByQuery(bson.M{"_id": userID, "password": auth.HashPassword(password)}))
}

func (auth *authRepo) GetByCredentials(request *models.AuthModel) (*models.AuthModel, error) {
	return auth.DecodeSingle(auth.col.GetSingleByQuery(bson.M{"phone": request.Phone, "password": auth.HashPassword(request.Password)}))
}

func (auth *authRepo) GetByPhone(phone *string) (*models.AuthModel, error) {
	return auth.DecodeSingle(auth.col.GetSingleByQuery(bson.M{"phone": phone}))
}

func (auth *authRepo) GetByID(id *primitive.ObjectID) (*models.AuthModel, error) {
	return auth.DecodeSingle(auth.col.GetSingleByQuery(bson.M{"_id": id}))
}

func (auth *authRepo) ReturnClient() db.StartMongoClient {
	return auth.client
}


// PRIVATE
func (auth *authRepo) DecodeSingle(dbResult *mongo.SingleResult) (*models.AuthModel, error) {
	auth.errRes = dbResult.Decode(&auth.authRes)
	if auth.errRes != nil {
		return nil, auth.errRes
	}
	return &auth.authRes, nil
}