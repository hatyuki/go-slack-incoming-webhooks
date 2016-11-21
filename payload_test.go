package slackIncomingWebhooks

import (
	"testing"
)

func TestPayload_ToJSON(t *testing.T) {
	payload := &Payload{Username: "test"}

	if _, err := payload.ToJSON(); err == nil {
		t.Error("should return an error")
	}

	payload.Text = "message"
	if got, err := payload.ToJSON(); err != nil {
		t.Errorf("returns an error: got %v", err)
	} else if want := `{"username":"test","text":"message"}`; got != want {
		t.Errorf("returns wrong JSON: got %v want %v", got, want)
	}
}
