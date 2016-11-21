package slackIncomingWebhooks

import (
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/pkg/errors"
)

const Version = "v0.0.2"

// Post messages to Slack
func Post(uri string, payload *Payload) error {
	json, err := payload.ToJSON()
	if err != nil {
		return err
	}

	response, err := http.PostForm(uri, url.Values{"payload": []string{json}})
	if err != nil {
		return errors.Wrap(err, "an error occurred when trying to send message")
	}

	if response.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(response.Body)
		defer response.Body.Close()
		return errors.Errorf("could not send message: %s", body)
	}

	return nil
}
