package mapClient

import (
	"fmt"
	"net/http"

	httpCode "angmorning.com/internal/libs/http/http-code"
	httpError "angmorning.com/internal/libs/http/http-error"
)

type naverMapClient struct {
	httpClient *http.Client
}

func newNaverClient() *naverMapClient {
	return &naverMapClient{
		httpClient: http.DefaultClient,
	}
}

func (naver *naverMapClient) Geocode(address string) (string, error) {

	req, err := http.NewRequest("GET", "https://naveropenapi.apigw.ntruss.com/map-geocode/v2/geocode", nil)
	if err != nil {
		return "", httpError.New(httpCode.InternalServerError, err.Error(), "")
	}

	req.Header.Add("x-ncp-apigw-api-key-id", "TEST")
	req.Header.Add("x-ncp-apigw-api-key", "TTTTTTTTTTTT")
	req.Header.Add("Accept", "application/json")

	res, err := naver.httpClient.Do(req)

	fmt.Println("!!!", res)
	return "naver", nil
}
