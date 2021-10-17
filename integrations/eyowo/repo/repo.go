package repo

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/klusters-core/api/integrations/eyowo"
	"github.com/klusters-core/api/integrations/eyowo/models"
	"github.com/klusters-core/api/utils"
	"fmt"
	"log"
	"net/http"
	"strings"
	"unicode/utf8"
)

var (
	mapDecoder utils.GeneralUtil
	NetWorker  = eyowo.NewNetworkService()
)

type (
	IDeveloperRepo interface {
		ValidateApp() (*models.AuthResponse, error)
		RefreshToken(refreshToken string) (*models.RefreshTokenResponse, error)
		SaveCheckoutBill(bill *models.CreateBillRequest) (*models.CreateBillResponse, error)
		VerifyBillPayment(paymentRef string) (*models.VerifyBillResponse, error)
		SendOTP(mobile string) (error)
		VerifyOTP(passCode string, mobile string) (*models.OtpResponse, error)
		MakePhoneTransfer(request *models.PhoneTransferRequest, token *string) (*models.PhoneTransferResponse, error)
	}

	developerRepo struct {
		XAppKey			string
		BaseUrl			string
		Mobile			string
	}
)

func NewDeveloperRepo(baseURL string, appKey string) *developerRepo {
	return &developerRepo{
		XAppKey: appKey,
		BaseUrl:   	baseURL,
	}
}

func (repo *developerRepo) ValidateApp() (*models.AuthResponse, error) {
	payload, err := json.Marshal(map[string]string{
		"mobile": repo.Mobile,
	})

	if err != nil {
		return nil, err
	}

	statusCode, res, err := eyowo.ConfigureWriteRequest("POST", fmt.Sprintf("%s/v1/users/auth/validate", repo.BaseUrl), bytes.NewBuffer(payload), eyowo.ConfigureAuthHeaders(repo.XAppKey))
	if err != nil {
		return nil, err
	}

	var result *models.AuthResponse
	if err := utils.TransformMAP(&res, &result); err != nil {
		log.Println(err.Error())
		return nil, errors.New("error while attempting to unmarshal auth result")
	}

	if statusCode != http.StatusOK {
		return nil, errors.New(result.Errors)
	}

	return result, nil
}

func (repo *developerRepo) RefreshToken(refreshToken string) (*models.RefreshTokenResponse, error) {
	payload, _ := json.Marshal(map[string]string{
		"refreshToken": refreshToken,
	})

	statusCode, res, err := eyowo.ConfigureWriteRequest("POST", fmt.Sprintf("%s/v1/users/accessToken", repo.BaseUrl), bytes.NewBuffer(payload), eyowo.ConfigureAuthHeaders(repo.XAppKey))
	if err != nil {
		log.Println(err.Error())
		return nil, errors.New("error encountered refreshing token")
	}

	log.Println(res)

	var result *models.RefreshTokenResponse
	_ = utils.TransformMAP(&res, &result)

	if statusCode != http.StatusOK {
		return nil, errors.New(result.Errors)
	}

	return result, nil
}

func (repo *developerRepo) SaveCheckoutBill(bill *models.CreateBillRequest) (*models.CreateBillResponse, error){
	payload := mapDecoder.ConvertStructToMap(*bill)
	statusCode, response, err := NetWorker.Post(fmt.Sprintf("%s/v1/checkout", repo.BaseUrl), eyowo.ConfigureAuthHeaders(repo.XAppKey), payload)
	if err != nil {
		log.Println(err)
		log.Println(errors.New("unable to create handshake with eyowo api"))
	}

	var checkOutBill *models.CreateBillResponse
	var logResponse interface{}
	err = json.Unmarshal(response, &checkOutBill)
	err = json.Unmarshal(response, &logResponse)
	if err != nil {
		return nil, errors.New("unable to decode checkout response")
	}

	log.Println(logResponse)

	if statusCode != http.StatusCreated {
		log.Println(checkOutBill)
		return nil, errors.New("unable to decode checkout success response")
	}

	return checkOutBill, nil
}

func (repo *developerRepo) VerifyBillPayment(paymentRef string) (*models.VerifyBillResponse, error) {
	_, res, err := NetWorker.Get(fmt.Sprintf("%s/v1/checkout/status/%s", repo.BaseUrl, paymentRef), eyowo.ConfigureAuthHeaders(repo.XAppKey))
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	var result *models.VerifyBillResponse
	err = json.Unmarshal(res, &result)
	if err != nil {
		return nil, errors.New("unable to decode checkout response")
	}

	return result, nil
}

func (repo *developerRepo) SendOTP(mobile string) (error) {
	_, i := utf8.DecodeRuneInString(mobile)

	payload, err := json.Marshal(map[string]string{
		"mobile": "234"+mobile[i:],
		"factor": "sms",
	})

	log.Println(mobile, "this is the payload")

	statusCode, res, err := NetWorker.Post(fmt.Sprintf("%s/v1/users/auth", repo.BaseUrl), eyowo.ConfigureAuthHeaders(repo.XAppKey), bytes.NewBuffer(payload))
	if err != nil {
		log.Println(err.Error())
		return err
	}

	if statusCode != http.StatusOK {
		var result models.ErrorResponse
		err = json.Unmarshal(res, &result)
		if err != nil {
			log.Println(err)
		}

		return errors.New(result.Error)
	}

	return nil
}

func (repo *developerRepo) VerifyOTP(passCode string, mobile string) (*models.OtpResponse, error) {
	_, i := utf8.DecodeRuneInString(mobile)

	payload, err := json.Marshal(map[string]string{
		"mobile": "234"+mobile[i:],
		"factor": "sms",
		"passcode": passCode,
	})

	statusCode, res, err := NetWorker.Post(fmt.Sprintf("%s/v1/users/auth", repo.BaseUrl), eyowo.ConfigureAuthHeaders(repo.XAppKey), bytes.NewBuffer(payload))
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	var result *models.OtpResponse
	err = json.Unmarshal(res, &result)
	if err != nil {
		return nil, errors.New("unable to decode verify OTP response")
	}

	if statusCode != http.StatusOK {
		var result models.ErrorResponse
		err = json.Unmarshal(res, &result)
		if err != nil {
			log.Println(err)
		}

		return nil, errors.New(result.Error)
	}

	return result, nil
}

func (repo *developerRepo) MakePhoneTransfer(request *models.PhoneTransferRequest, token *string) (*models.PhoneTransferResponse, error) {
	newBuffer := new(bytes.Buffer)
	err := json.NewEncoder(newBuffer).Encode(*request)
	headers := map[string]string{
		"X-App-Key": repo.XAppKey,
		"X-App-Wallet-Access-Token":  strings.Trim(*token, "\""),
		"Content-Type": "application/json",
		"User-Agent":   "Mozilla/5.0 (Unknown; Linux) AppleWebKit/538.1 (KHTML, like Gecko) Chrome/v1.0.0 Safari/538.1",
	}

	statusCode, response, err := eyowo.ConfigureWriteRequest("POST", fmt.Sprintf("%s/v1/users/transfers/phone", repo.BaseUrl), newBuffer, headers)
	if err != nil {
		log.Println(err)
		log.Println(errors.New("unable to create handshake with eyowo api"))
	}

	log.Println(statusCode, response, "status code from eyowo")

	var res *models.PhoneTransferResponse
	err = utils.TransformMAP(&response, &res)
	if err != nil {
		log.Println(err)
		return nil, errors.New("unable to decode checkout response")
	}

	if statusCode != http.StatusOK {
		return nil, errors.New("unable to decode transfer success response")
	}

	return res, nil
}