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

func (c *HackHourSlackClient) GetHackHourClient() (*api.HackHourClient) {
	return c.hhClient
}
func (c *HackHourSlackClient) SendToSessionThread(msg string, url string) (channel string, ts string, err error) {
	session, err := c.hhClient.GetSession()
	if err != nil {
		fmt.Println(err)
		return
	}
	blocks := []slack.Block{
		&slack.SectionBlock{
			Type: slack.MBTSection,
			Text: slack.NewTextBlockObject(slack.MarkdownType, msg, false, false),
			Accessory: &slack.Accessory{
				ButtonElement: &slack.ButtonBlockElement{
					Type: slack.METButton,
					Text: slack.NewTextBlockObject(slack.PlainTextType, "Open in Github", false, false),
					URL:  url,
				},
			},
		},
	}
	if url != "" {
		_, _, err = c.slack.PostMessage("C06SBHMQU8G", slack.MsgOptionText(msg, false), slack.MsgOptionBlocks(blocks...), slack.MsgOptionTS(session.MessageTs))
	} else {
		channel, ts, err = c.slack.PostMessage("C06SBHMQU8G", slack.MsgOptionText(msg, false), slack.MsgOptionTS(session.MessageTs))
	}

	return
}
