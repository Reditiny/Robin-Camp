package client

import (
	"fmt"
	"net/http"
	"time"

	"assignment/config"
	"assignment/openapi"
)

type BoxOfficeClient struct {
	maxRetry int
	cli      http.Client
}

func NewBoxOfficeClient() *BoxOfficeClient {
	return &BoxOfficeClient{
		cli: http.Client{
			Timeout: 3 * time.Second,
			Transport: &http.Transport{
				MaxIdleConns:    100,
				IdleConnTimeout: 90 * time.Second,
			},
		},
		maxRetry: 3,
	}
}

var _ openapi.HttpRequestDoer = &BoxOfficeClient{}

func (boc *BoxOfficeClient) Do(req *http.Request) (*http.Response, error) {

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("apifoxToken", config.Conf.BoxOfficeApiKey)
	req.Header.Set("User-Agent", "MyApp/1.0")

	for i := 1; i < boc.maxRetry; i++ {
		response, err := boc.cli.Do(req)
		if err != nil {
			fmt.Printf("retry %d BoxOfficeApiKey\n", i)
			continue
		}
		return response, nil
	}

	return boc.cli.Do(req)
}
