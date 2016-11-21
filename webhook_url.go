package slackIncomingWebhooks

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/goulash/xdg"
	"github.com/pkg/errors"
)

const configfile = "slack-incoming-webhooks/url"

func ReadConfig(path string) (string, error) {
	if path == "" {
		if path := xdg.FindConfig(configfile); path == "" {
			return "", errors.New("no webhook URL given: use `--webhook-url` or `--configure` option")
		} else {
			return readConfig(path)
		}
	} else if strings.Index(path, "http://") == 0 || strings.Index(path, "https://") == 0 {
		return path, nil
	} else {
		return readConfig(path)
	}
}

func WriteConfig(uri string) error {
	if path := xdg.UserConfig(configfile); path == "" {
		return errors.New("could not create config file")
	} else {
		return writeConfig(path, uri)
	}
}

func readConfig(path string) (string, error) {
	fp, err := os.Open(path)

	if err != nil {
		return "", errors.Wrap(err, "could not open config file")
	}
	defer fp.Close()

	buffer, err := ioutil.ReadAll(fp)
	if err != nil {
		return "", errors.Wrap(err, "could not read config file")
	} else if url := string(buffer); url == "" {
		return "", errors.Errorf("no webhook URL given: %s", path)
	} else {
		return url, nil
	}
}

func writeConfig(path string, uri string) error {
	dir := filepath.Dir(path)
	if _, err := os.Stat(dir); err != nil {
		if err := os.MkdirAll(dir, 0700); err != nil {
			return errors.Wrap(err, "could not create config directory")
		}
	}

	return ioutil.WriteFile(path, []byte(uri), 0600)
}
