package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type HackHourClient struct {
	client  *http.Client
	apiKey  string
	baseURL string
}

func NewHackHourClient(apiKey string) *HackHourClient {
	return &HackHourClient{
		client:  &http.Client{},
		apiKey:  apiKey,
		baseURL: "https://hackhour.hackclub.com",
	}
}

func (c *HackHourClient) createRequest(method string, endpoint string) (*http.Request, error) {
	if endpoint[0] == '/' {
		return nil, fmt.Errorf("endpoint should not have a leading slash")
	}
	req, err := http.NewRequest(method, fmt.Sprintf("%v/%v", c.baseURL, endpoint), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", c.apiKey))
	return req, nil
}

func (c *HackHourClient) Ping() bool {
	req, err := c.createRequest("GET", "ping")
	if err != nil {
		return false
	}
	_, err = c.client.Do(req)
	return err == nil
}

type StatusResponse struct {
	ActiveSessions    int  `json:"activeSessions"`
	AirtableConnected bool `json:"airtableConnected"`
	SlackConnected    bool `json:"slackConnected"`
}

func (c *HackHourClient) Status() (*StatusResponse, error) {
	req, err := c.createRequest("GET", "status")
	if err != nil {
		return nil, err
	}
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	out := &StatusResponse{}
	err = json.NewDecoder(resp.Body).Decode(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}
