package slackIncomingWebhooks

import(
	"testing"
)

func TestEmptyMessage (t *testing.T) {
	payload := &Payload{Username: "test"}

	if _, err := payload.ToJSON(); err == nil {
		t.Error("does not return error")
	}
}
