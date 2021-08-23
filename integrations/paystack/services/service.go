//=============================================================================
// developer: boxlesslabsng@gmail.com
// requires a valid pay stack API key to be passed to the repo layer
// exposes the pay stack services for clients
//=============================================================================

/**
 **
 * @struct payStackService
 * @interface PayStackService
 **
**/

package services

import (
	"bytes"
	"encoding/json"
	"github.com/klusters-core/api/integrations/paystack/models"
	"github.com/klusters-core/api/integrations/paystack/repo"
	"github.com/klusters-core/api/utils"
)

func newPayStackService(APIKey string) *payStackService {
	payStack := repo.NewPayStackRepo(APIKey)
	return &payStackService{payStack}
}

type (
	payStackService struct {
		repo.PayStackRepo
	}

	PayStackService interface {
		VerifyTransaction(reference string) (*models.VerifyTransaction, error)
		ListTransactions() (*models.ListTransactions, error)
		GetTransaction(id string) (*models.FetchTransaction, error)
		InitiateTransaction(reference string, amount int, email string) (*models.InitializeTransaction, error)
		ChargeAuthorization(authCode string, amount int, email string) (*models.ChargeAuthorization, error)
	}
)

func (ps *payStackService) VerifyTransaction(reference string) (*models.VerifyTransaction, error) {
	ps.SetHTTPMethod("GET")
	ps.SetURL("transaction/verify/"+reference)
	res, err := ps.SendRequest()
	if err != nil {
		return nil, err
	}

	var result *models.VerifyTransaction
	if err := utils.TransformMAP(&res, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (ps *payStackService) ListTransactions() (*models.ListTransactions, error) {
	ps.SetHTTPMethod("GET")
	ps.SetURL("transaction")
	res, err := ps.SendRequest()
	if err != nil {
		return nil, err
	}

	var result *models.ListTransactions
	if err := utils.TransformMAP(&res, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (ps *payStackService) GetTransaction(id string) (*models.FetchTransaction, error) {
	ps.SetHTTPMethod("GET")
	ps.SetURL("transaction/"+id)
	res, err := ps.SendRequest()
	if err != nil {
		return nil, err
	}

	var result *models.FetchTransaction
	if err := utils.TransformMAP(&res, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (ps *payStackService) InitiateTransaction(reference string, amount int, email string) (*models.InitializeTransaction, error) {
	payload, err := json.Marshal(map[string]interface{} {
		"reference": reference,
		"amount": amount,
		"email": email,
	})
	if err != nil {
		return nil, err
	}

	ps.SetHTTPMethod("POST")
	ps.SetURL("transaction/initialize")
	ps.SetPayload(bytes.NewBuffer(payload))
	res, err := ps.SendRequest()
	if err != nil {
		return nil, err
	}

	var result *models.InitializeTransaction
	_ = utils.TransformMAP(&res, &result)

	return result, nil
}

func (ps *payStackService) ChargeAuthorization(authCode string, amount int, email string) (*models.ChargeAuthorization, error) {
	payload, err := json.Marshal(map[string]interface{} {
		"reference": authCode,
		"amount": amount,
		"email": email,
	})
	if err != nil {
		return nil, err
	}

	ps.SetHTTPMethod("POST")
	ps.SetURL("transaction/charge_authorization")
	ps.SetPayload(bytes.NewBuffer(payload))
	res, err := ps.SendRequest()
	if err != nil {
		return nil, err
	}

	var result *models.ChargeAuthorization
	_ = utils.TransformMAP(&res, &result)

	return result, nil
}