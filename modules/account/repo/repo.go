//=============================================================================
// developer: boxlesslabsng@gmail.com
//=============================================================================

/**
 * Defines an abstract class for accounts
 * Provides an interfaces for accessing accountRepo type
 **
 * @struct accountRepo
**/


package repo

import (
	"errors"
	"github.com/klusters-core/api/config/db"
	"github.com/klusters-core/api/config/error_response"
	model "github.com/klusters-core/api/modules/account/models"
	"github.com/klusters-core/api/modules/auth/models"
	"github.com/klusters-core/api/modules/auth/repo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

type (
	accountRepo struct {
		col 	db.MongoInterface
		client 	db.StartMongoClient
	}

	AccountRepo interface {
		CreateAccount(req *model.AccountRequest, auth *models.AuthModel) (*model.AccountsModel, error)
		GetAccount(id primitive.ObjectID) (*model.AccountsModel, error)
		GetByPhone(phone string) (*model.AccountsModel, error)
	}
)

func NewAccountRepo(client db.StartMongoClient) *accountRepo {
	col := db.NewMongoCollection("account", client)
	return &accountRepo{col: col, client:client}
}

func (account *accountRepo) CreateAccount(req *model.AccountRequest, auth *models.AuthModel) (*model.AccountsModel, error) {
	_, err := account.GetByPhone(req.Phone)
	if err == nil {
		return nil, errors.New(error_response.DuplicateError{Resource: "user account"}.Error())
	}

	Auth := repo.NewAuthRepo(account.ReturnClient())
	_, err = Auth.Create(auth)


	newAccount := model.SetAccount(req)
	newAccount.NewID()
	newAccount.MakeOwner()
	result, err := account.col.AddSingle(newAccount)
	if err != nil {
		return nil, errors.New(error_response.NotCreated{Resource: "user account"}.Error())
	}
	return account.GetAccount(result.DocID)
}

func (account *accountRepo) GetAccount(id primitive.ObjectID) (*model.AccountsModel, error) {
	return account.DecodeSingle(account.col.GetSingleById(id))
}

func (account *accountRepo) GetByPhone(phone string) (*model.AccountsModel, error) {
	return account.DecodeSingle(account.col.GetSingleByQuery(bson.M{"phone": phone}))
}

func (account *accountRepo) DecodeSingle(dbResult *mongo.SingleResult) (*model.AccountsModel, error) {
	var user *model.AccountsModel
	decodeErr := dbResult.Decode(&user)
	if decodeErr != nil {
		log.Println(decodeErr)
		return nil, decodeErr
	}
	return user, nil
}

func (account *accountRepo) ReturnClient() db.StartMongoClient {
	return account.client
}