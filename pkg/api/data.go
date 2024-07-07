package api

import (
	"encoding/json"
	"fmt"
	"time"
)

func get[T interface{}](c *HackHourClient, endpoint string) (*T, error) {
	req, err := c.createGetRequest(fmt.Sprintf("api/%v/%v", endpoint, c.slackId))
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
		Data  T      `json:"data"`
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

// GET /api/session/:slackId
type Session struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	Time      int       `json:"time"`
	Elapsed   int       `json:"elapsed"`
	Remaining int       `json:"remaining"`
	EndTime   time.Time `json:"endTime"`
	Goal      string    `json:"goal"`
	Paused    bool      `json:"paused"`
	Completed bool      `json:"completed"`
	MessageTs string    `json:"messageTs"`
}

func (c *HackHourClient) GetSession() (*Session, error) {
	return get[Session](c, "session")
}

// GET /api/stats/:slackId
type Stats struct {
	Sessions int `json:"sessions"`
	Total    int `json:"total"`
}

func (c *HackHourClient) GetStats() (*Stats, error) {
	return get[Stats](c, "stats")
}

// GET /api/goals/:slackId
type Goals []struct {
	Name    string `json:"name"`
	Minutes int    `json:"minutes"`
}

func (c *HackHourClient) GetGoals() (*Goals, error) {
	return get[Goals](c, "goals")
}

// GET /api/history/:slackId
type History []struct {
	CreatedAt time.Time `json:"createdAt"`
	Time      int       `json:"time"`
	Elapsed   int       `json:"elapsed"`
	Goal      string    `json:"goal"`
	Ended     bool      `json:"ended"`
	Work      string    `json:"work"`
}

func (c *HackHourClient) GetHistory() (*History, error) {
	return get[History](c, "history")
}
