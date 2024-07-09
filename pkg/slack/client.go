package slack

import (
	"fmt"

	"github.com/rutmanz/hackhour-go/pkg/api"
	"github.com/slack-go/slack"
)

type HackHourSlackClient struct {
	slack    *slack.Client
	hhClient *api.HackHourClient
}

func CreateClient(hhClient *api.HackHourClient, token string) *HackHourSlackClient {
	api := slack.New(token)
	return &HackHourSlackClient{
		hhClient: hhClient,
		slack:    api,
	}
}
func (c *HackHourSlackClient) CheckAuth() (*slack.AuthTestResponse, error) {
	return c.slack.AuthTest()

}
func (c *HackHourSlackClient) SendToSessionThread(msg string) (channel string, ts string, err error) {
	session, err := c.hhClient.GetSession()
	if err != nil {
		fmt.Println(err)
		return
	}
	channel, ts, err = c.slack.PostMessage("C06SBHMQU8G", slack.MsgOptionText(msg, false), slack.MsgOptionTS(session.MessageTs))
	return
}
