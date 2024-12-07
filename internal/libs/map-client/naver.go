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

	req, err := http.NewRequest("GET", fmt.Sprintf("https://naveropenapi.apigw.ntruss.com/map-geocode/v2/geocode?query=%s", address), nil)
	if err != nil {
		return "", httpError.New(httpCode.InternalServerError, err.Error(), "")
	}

	req.Header.Add("x-ncp-apigw-api-key-id", "x8ispzy2ng")
	req.Header.Add("x-ncp-apigw-api-key", "zzxo6tuI30z08OvfIDDtfVqcRbvzronXWqFj1kJc")
	req.Header.Add("Accept", "application/json")

	res, err := naver.httpClient.Do(req)

	fmt.Println("!!!", res)
	return "naver", nil
}
