package slackIncomingWebhooks

import (
	"testing"
)

func TestPost(t *testing.T) {
	uri := "https://example.com"
	payload := &Payload{}

	if err := Post(uri, payload); err != err {
		t.Error("should return an error")
	}
}
