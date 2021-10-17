package repo

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/klusters-core/api/integrations/eyowo/models"
	"github.com/klusters-core/api/utils"
	"log"
	"net/http"
	"strings"
	"unicode/utf8"
)

var (
	NetWorker  = utils.NewNetworkService()
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


// public
func (repo *developerRepo) ValidateApp() (*models.AuthResponse, error) {
	payload, _ := json.Marshal(map[string]string{
		"mobile": repo.Mobile,
	})

	statusCode, res, err := NetWorker.Post(fmt.Sprintf("%s/v1/users/auth/validate", repo.BaseUrl), repo.ConfigureAuthHeaders(repo.XAppKey), bytes.NewBuffer(payload))
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	var result *models.AuthResponse
	if err = json.Unmarshal(res, &result); err != nil {
		return nil, errors.New("error while attempting to unmarshal auth result")
	}

	if statusCode != http.StatusOK {
		return nil, errors.New(result.Error)
	}

	return result, nil
}

func (repo *developerRepo) RefreshToken(refreshToken string) (*models.RefreshTokenResponse, error) {
	payload, _ := json.Marshal(map[string]string{
		"refreshToken": refreshToken,
	})

	statusCode, res, err := NetWorker.Post(fmt.Sprintf("%s/v1/users/accessToken", repo.BaseUrl), repo.ConfigureAuthHeaders(repo.XAppKey), bytes.NewBuffer(payload))
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	var result *models.RefreshTokenResponse
	if err = json.Unmarshal(res, &result); err != nil {
		return nil, errors.New("error encountered refreshing token")
	}

	if statusCode != http.StatusOK {
		return nil, errors.New(result.Error)
	}

	return result, nil
}

func (repo *developerRepo) SaveCheckoutBill(bill *models.CreateBillRequest) (*models.CreateBillResponse, error){
	payload, _ := json.Marshal(bill)
	statusCode, res, err := NetWorker.Post(fmt.Sprintf("%s/v1/checkout", repo.BaseUrl), repo.ConfigureAuthHeaders(repo.XAppKey), bytes.NewBuffer(payload))
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	var result *models.CreateBillResponse
	if err = json.Unmarshal(res, &result); err != nil {
		return nil, errors.New("unable to decode checkout response")
	}

	if statusCode != http.StatusOK {
		return nil, errors.New(result.Error)
	}

	return result, nil
}

func (repo *developerRepo) VerifyBillPayment(paymentRef string) (*models.VerifyBillResponse, error) {
	statusCode, res, err := NetWorker.Get(fmt.Sprintf("%s/v1/checkout/status/%s", repo.BaseUrl, paymentRef), repo.ConfigureAuthHeaders(repo.XAppKey))
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	var result *models.VerifyBillResponse
	if err = json.Unmarshal(res, &result); err != nil {
		return nil, errors.New("unable to decode checkout response")
	}

	if statusCode != http.StatusOK {
		return nil, errors.New(result.Error)
	}

	return result, nil
}

func (repo *developerRepo) SendOTP(mobile string) (error) {
	_, i := utf8.DecodeRuneInString(mobile)

	payload, err := json.Marshal(map[string]string{
		"mobile": "234"+mobile[i:],
		"factor": "sms",
	})

	statusCode, res, err := NetWorker.Post(fmt.Sprintf("%s/v1/users/auth", repo.BaseUrl), repo.ConfigureAuthHeaders(repo.XAppKey), bytes.NewBuffer(payload))
	if err != nil {
		log.Println(err.Error())
		return err
	}

	var result *models.OtpResponse
	if err = json.Unmarshal(res, &result); err != nil {
		return errors.New("unable to send OTP at this time")
	}

	if statusCode != http.StatusOK {
		return errors.New(result.Error)
	}

	return nil
}

func (repo *developerRepo) VerifyOTP(passCode string, mobile string) (*models.OtpResponse, error) {
	_, i := utf8.DecodeRuneInString(mobile)

	payload, _ := json.Marshal(map[string]string{
		"mobile": "234"+mobile[i:],
		"factor": "sms",
		"passcode": passCode,
	})

	statusCode, res, err := NetWorker.Post(fmt.Sprintf("%s/v1/users/auth", repo.BaseUrl), repo.ConfigureAuthHeaders(repo.XAppKey), bytes.NewBuffer(payload))
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	var result *models.OtpResponse
	if err = json.Unmarshal(res, &result); err != nil {
		return nil, errors.New("unable to decode verify OTP response")
	}

	if statusCode != http.StatusOK {
		return nil, errors.New(result.Error)
	}

	return result, nil
}

func (repo *developerRepo) MakePhoneTransfer(request *models.PhoneTransferRequest, token *string) (*models.PhoneTransferResponse, error) {
	payload, _ := json.Marshal(request)
	statusCode, res, err := NetWorker.Post(fmt.Sprintf("%s/v1/users/transfers/phone", repo.BaseUrl), repo.ConfigureAuthHeaders2(token), bytes.NewBuffer(payload))
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	var result *models.PhoneTransferResponse
	if err = json.Unmarshal(res, &result); err != nil {
		return nil, errors.New("unable to decode phone transfer response")
	}

	if statusCode != http.StatusOK {
		return nil, errors.New(result.Error)
	}

	return result, nil
}


// private
func (repo *developerRepo) ConfigureAuthHeaders(appKey string) map[string]string {
	return map[string]string{
		"X-App-Key":    appKey,
		"Content-Type": "application/json",
		"User-Agent":   "Mozilla/5.0 (Unknown; Linux) AppleWebKit/538.1 (KHTML, like Gecko) Chrome/v1.0.0 Safari/538.1",
	}
}

func (repo *developerRepo) ConfigureAuthHeaders2(token *string) map[string]string {
	return map[string]string{
		"X-App-Key":                 repo.XAppKey,
		"X-App-Wallet-Access-Token": strings.Trim(*token, "\""),
		"Content-Type":              "application/json",
		"User-Agent":                "Mozilla/5.0 (Unknown; Linux) AppleWebKit/538.1 (KHTML, like Gecko) Chrome/v1.0.0 Safari/538.1",
	}
}