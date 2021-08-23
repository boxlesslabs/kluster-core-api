//=============================================================================
// developer: boxlesslabsng@gmail.com
// prepares transport for interfacing with pay stack apis
// requires a valid pay stack API key for implementation
//=============================================================================

/**
 **
 * @struct payStackRepo
 * @interface PayStackRepo
 **
**/

package repo

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"time"
)

func NewPayStackRepo(key string) *payStackRepo {
	return &payStackRepo{
		Key:        key,
	}
}

type (
	payStackRepo struct {
		Key				string
		HTTPMethod 		string
		URL				string
		Payload 		io.Reader
	}

	PayStackRepo interface {
		SetHTTPMethod(method string)
		SetURL(url string)
		SendRequest() (interface{}, error)
		SetPayload(payload io.Reader)
	}
)

func (repo *payStackRepo) SetHTTPMethod(method string) {
	repo.HTTPMethod = method
}

func (repo *payStackRepo) SetURL(url string) {
	repo.URL = url
}

func (repo *payStackRepo) SetPayload(payload io.Reader) {
	repo.Payload = payload
}

func (repo *payStackRepo) SendRequest() (interface{}, error) {
	if repo.HTTPMethod == "" {
		return nil, errors.New("HTTP Method is required")
	}

	if repo.URL == "" {
		return nil, errors.New("API URL is required")
	}

	client := repo.SetClient()
	var buf io.Reader

	if repo.HTTPMethod == "POST" {
		buf = repo.Payload
	}else {
		buf = nil
	}

	request, err := http.NewRequest(repo.HTTPMethod, repo.URL, buf)
	if err != nil {
		return nil, err
	}

	for  k, v := range repo.SetJSONHeader() {
		request.Header.Set(k, v)
	}

	response, err := client.Do(request)
	if err != nil {
		//log.Fatal("Error occurred dispatching request", err.Error())
		return nil, err
	}

	err = response.Body.Close()
	if err != nil {
		return nil, err
	}

	var res map[string]interface{}
	_ = json.NewDecoder(response.Body).Decode(&res)

	return res, nil
}

func (repo *payStackRepo) SetJSONHeader() map[string]string {
	return map[string]string {
		"Authorization": "Bearer " + repo.Key,
		"Content-Type":	"application/json",
	}
}

func (repo *payStackRepo) SetClient() http.Client {
	timeout := time.Duration(60 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}

	return client
}