package slackIncomingWebhooks

import (
	"encoding/json"

	"github.com/pkg/errors"
)

type Payload struct {
	Channel   string `json:"channel,omitempty"`
	Username  string `json:"username,omitempty"`
	IconEmoji string `json:"icon_emoji,omitempty"`
	IconURL   string `json:"icon_url,omitempty"`
	Text      string `json:"text,omitempty"`
}

func (payload *Payload) ToJSON() (string, error) {
	if payload.Text == "" {
		return "", errors.New("message body should not be empty")
	}

	bytes, err := json.Marshal(payload)
	if err != nil {
		return "", errors.Wrap(err, "failed to convert json")
	}

	return string(bytes), nil
}
