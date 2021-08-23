package repo

import (
	"github.com/klusters-core/api/config/db"
	"github.com/klusters-core/api/modules/auth/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type (
	authRepo struct {
		col 	db.MongoInterface
		client 	db.StartMongoClient
	}

	AuthRepo interface {
		Create(req *models.AuthRequest) (primitive.ObjectID, error)
		GetByCredentials(req *models.AuthRequest) (*models.AuthModel, error)
		GetByPhone(phone string) (*models.AuthModel, error)
		ReturnClient() db.StartMongoClient
		UpdatePassword(id primitive.ObjectID, password string) (*models.AuthModel, error)
		ComparePasswords(userID primitive.ObjectID, password string) (*models.AuthModel, error)
	}
)

func NewAuthRepo(client db.StartMongoClient) *authRepo {
	col := db.NewMongoCollection("auth", client)
	return &authRepo{col: col, client:client}
}

func (auth *authRepo) Create(req *models.AuthRequest) (primitive.ObjectID, error) {
	newAuth := models.SetAuth(req)
	newAuth.NewID()
	newAuth.TimeStamp()
	newAuth.EncryptPassword()

	result, err := auth.col.AddSingle(newAuth)
	if err != nil {
		return primitive.ObjectID{}, err
	}
	return result.DocID, nil
}

func (auth *authRepo) UpdatePassword(id primitive.ObjectID, password string) (*models.AuthModel, error) {
	result, err := auth.col.UpdateById(id, models.UpdatePasswordQuery(password))
	if err != nil {
		return nil, err
	}

	return auth.DecodeSingle(result)
}

func (auth *authRepo) ComparePasswords(userID primitive.ObjectID, password string) (*models.AuthModel, error) {
	return auth.DecodeSingle(auth.col.GetSingleByProjection(models.GetPasswordQuery(userID, password), bson.M{"password": 1}))
}

func (auth *authRepo) GetByCredentials(req *models.AuthRequest) (*models.AuthModel, error) {
	return auth.DecodeSingle(auth.col.GetSingleByQuery(models.GetByCredentialsQuery(req.Phone, req.Password)))
}

func (auth *authRepo) GetByPhone(phone string) (*models.AuthModel, error) {
	return auth.DecodeSingle(auth.col.GetSingleByQuery(models.GetByPhoneQuery(phone)))
}

func (auth *authRepo) ReturnClient() db.StartMongoClient {
	return auth.client
}

func (auth *authRepo) DecodeSingle(dbResult *mongo.SingleResult) (*models.AuthModel, error) {
	var user *models.AuthModel
	decodeErr := dbResult.Decode(&user)
	if decodeErr != nil {
		return nil, decodeErr
	}
	return user, nil
}