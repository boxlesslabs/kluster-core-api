package utils

import (
	"github.com/go-resty/resty/v2"
	"log"
)

type networkService struct {
}

type INetwork interface {
	Get(url string, headers map[string]string) (int, []byte, error)
	Delete(url string, headers map[string]string) (int, []byte, error)
	Post(url string, headers map[string]string, body interface{})  (int, []byte, error)
	Patch(url string, headers map[string]string, body interface{})  (int, []byte, error)
}

func NewNetworkService() *networkService {
	return &networkService{}
}

func (service *networkService) Get(url string, headers map[string]string) (int,[]byte,error) {
	client := resty.New()

	request := client.R()

	if headers != nil {
		request = request.SetHeaders(headers)
	}

	response, err := request.Get(url)
	if err != nil {
		log.Println(err)
	}

	return response.StatusCode(), response.Body(), err
}

func (service *networkService) Post(url string, headers map[string]string, body interface{})  (int, []byte, error) {
	client := resty.New()

	request := client.R()

	if headers != nil {
		request = request.SetHeaders(headers)
	}

	if body != nil {
		request = request.SetBody(body)
	}

	response, err := request.Post(url)
	if err != nil {
		log.Println(err)
	}

	return response.StatusCode(),response.Body(),err
}

func (service *networkService) Patch(url string, headers map[string]string, body interface{})  (int, []byte, error) {
	client := resty.New()

	request := client.R()

	if headers != nil {
		request = request.SetHeaders(headers)
	}

	if body != nil {
		request = request.SetBody(body)
	}

	response, err := request.Patch(url)
	if err != nil {
		log.Println(err)
	}

	return response.StatusCode(),response.Body(),err
}

func (service *networkService) Delete(url string, headers map[string]string) (int,[]byte,error) {
	client := resty.New()

	request := client.R()

	if headers != nil {
		request = request.SetHeaders(headers)
	}

	response, err := request.Delete(url)
	if err != nil {
		log.Println(err)
	}

	return response.StatusCode(), response.Body(), err
}

