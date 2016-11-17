package slackIncomingWebhooks

import (
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/pkg/errors"
)

const Version = "v0.0.1"

// Client object
type Client struct {
	webhookURL string
}

// NewClient creates a new Client object.
func NewClient(webhookURL string) (client *Client, err error) {
	uri, err := GetWebhookURL(webhookURL)

	if err == nil {
		client = &Client{webhookURL: uri}
	}

	return
}

// Post message into Slack
func (client *Client) Post(payload *Payload) (err error) {
	json, err := payload.ToJSON()
	if err != nil {
		return
	}

	response, err := http.PostForm(client.webhookURL, url.Values{"payload": []string{json}})
	if err != nil {
		return
	}

	if response.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(response.Body)
		defer response.Body.Close()
		err = errors.Errorf("request failed: %s", body)
	}

	return
}
