package api

import "encoding/json"

func (c *HackHourClient) Ping() bool {
	req, err := c.createGetRequest("ping")
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
	req, err := c.createGetRequest("status")
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
