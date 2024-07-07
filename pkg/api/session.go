package api

import (
	"encoding/json"
	"fmt"
)

func post[Req interface{}, Resp interface{}](c *HackHourClient, endpoint string, body *Req) (*Resp, error) {
	b, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	req, err := c.createPostRequest(fmt.Sprintf("api/%v/%v", endpoint, c.slackId), b)
	if err != nil {
		return nil, err
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	out := &struct {
		Ok    bool   `json:"ok"`
		Error string `json:"error"`
		Data  Resp   `json:"data"`
	}{}

	err = json.NewDecoder(resp.Body).Decode(out)
	if err != nil {
		return nil, err
	}
	if !out.Ok {
		return nil, fmt.Errorf("API Error: %v", out.Error)
	}

	return &out.Data, nil
}

// POST /api/start/:slackId
type SessionStartRequest struct {
	Work string `json:"work"`
}
type SessionStartResponse struct {
	ID        string `json:"id"`
	SlackID   string `json:"slackId"`
	CreatedAt string `json:"createdAt"`
}

func (c *HackHourClient) SessionStart(activity string) (*SessionStartResponse, error) {
	return post[SessionStartRequest, SessionStartResponse](c, "start", &SessionStartRequest{Work: activity})
}

// POST /api/pause/:slackId

type SessionPauseResponse struct {
	ID        string `json:"id"`
	SlackID   string `json:"slackId"`
	CreatedAt string `json:"createdAt"`
	Paused    bool   `json:"paused"`
}

func (c *HackHourClient) SessionPause() (*SessionPauseResponse, error) {
	return post[struct{}, SessionPauseResponse](c, "pause", &struct{}{})
}

// POST /api/cancel/:slackId
type SessionCancelResponse struct {
	ID        string `json:"id"`
	SlackID   string `json:"slackId"`
	CreatedAt string `json:"createdAt"`
}

func (c *HackHourClient) SessionCancel(activity string) (*SessionCancelResponse, error) {
	return post[struct{}, SessionCancelResponse](c, "cancel", &struct{}{})
}
