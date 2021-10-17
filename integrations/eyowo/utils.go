package eyowo

import (
	"encoding/json"
	"io"
	"net/http"
	"time"

	secrets "github.com/klusters-core/api/config/secrets"
)

func ConfigureAuthHeaders(appKey string) map[string]string {
	return map[string]string{
		"X-App-Key":    appKey,
		"Content-Type": "application/json",
		"User-Agent":   "Mozilla/5.0 (Unknown; Linux) AppleWebKit/538.1 (KHTML, like Gecko) Chrome/v1.0.0 Safari/538.1",
	}
}

func ConfigureAuthHeaders2(token string) map[string]string {
	token = secrets.GetSecrets().EyowoRefreshToken

	return map[string]string{
		"X-App-Key":                 secrets.GetSecrets().EyowoAppKey,
		"X-App-Wallet-Access-Token": token,
		"Content-Type":              "application/json",
		"User-Agent":                "Mozilla/5.0 (Unknown; Linux) AppleWebKit/538.1 (KHTML, like Gecko) Chrome/v1.0.0 Safari/538.1",
	}
}

func ConfigureEyowoClient() http.Client {
	timeout := time.Duration(60 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}
	return client
}

func ConfigureWriteRequest(method string, url string, payload io.Reader, headers map[string]string) (int, interface{}, error) {
	client := ConfigureEyowoClient()
	var buf io.Reader

	if method == "POST" {
		buf = payload
	} else {
		buf = nil
	}

	request, err := http.NewRequest(method, url, buf)

	for k, v := range headers {
		request.Header.Set(k, v)
	}

	if err != nil {
		return 0, nil, err
	}

	response, err := client.Do(request)
	if err != nil {
		return 0, nil, err
	}

	defer response.Body.Close()
	var res interface{}
	_ = json.NewDecoder(response.Body).Decode(&res)

	return response.StatusCode, res, nil
}
