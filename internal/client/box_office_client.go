package client

import (
	"net/http"

	"assignment/config"
	"assignment/openapi"
)

type BoxOfficeClient struct {
	baseUrl string
	apiKey  string
	cli     http.Client
}

func NewBoxOfficeClient() *BoxOfficeClient {
	return &BoxOfficeClient{}
}

var _ openapi.HttpRequestDoer = &BoxOfficeClient{}

func (boc *BoxOfficeClient) Do(req *http.Request) (*http.Response, error) {

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("apifoxToken", config.Conf.BoxOfficeApiKey)
	req.Header.Set("User-Agent", "MyApp/1.0")

	return boc.cli.Do(req)
}
