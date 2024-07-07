package api

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

type HackHourClient struct {
	client  *http.Client
	apiKey  string
	baseURL string
	slackId string
}

func NewHackHourClient(apiKey string) *HackHourClient {
	return &HackHourClient{
		client:  &http.Client{},
		apiKey:  apiKey,
		baseURL: "https://hackhour.hackclub.com",
		slackId: "unused",
	}
}

func (c *HackHourClient) createGetRequest(endpoint string) (*http.Request, error) {
	return c.createRequestWithBody("GET", endpoint, nil)
}

func (c *HackHourClient) createPostRequest(endpoint string, body []byte) (*http.Request, error) {
	return c.createRequestWithBody("POST", endpoint, bytes.NewBuffer(body))
}

func (c *HackHourClient) createRequestWithBody(method string, endpoint string, body io.Reader) (*http.Request, error) {
	if endpoint[0] == '/' {
		return nil, fmt.Errorf("endpoint should not have a leading slash")
	}
	req, err := http.NewRequest(method, fmt.Sprintf("%v/%v", c.baseURL, endpoint), body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", c.apiKey))
	return req, nil
}
