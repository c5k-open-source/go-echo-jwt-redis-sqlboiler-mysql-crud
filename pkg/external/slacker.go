package external

import (
	"github.com/bluele/slack"
	"github.com/c5k-open-source/go-echo-jwt-redis-sqlboiler-mysql-crud/pkg/config"
)

type slackNotifier struct {
	enabled bool
	webHook *slack.WebHook
	channel string
}

func NewSlackNotifier() *slackNotifier {

	return &slackNotifier{
		enabled: config.Config.SlackEnabled,
		webHook: slack.NewWebHook(config.Config.SlackMainWebHook),
		channel: config.Config.SlackMainChannel,
	}
}

func (s *slackNotifier) Notify(text string) {
	if s.enabled {
		s.webHook.PostMessage(&slack.WebHookPostPayload{
			Text:    text,
			Channel: s.channel,
		})
	}
}
