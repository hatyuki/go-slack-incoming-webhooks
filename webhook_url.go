package slackIncomingWebhooks

import (
	"bufio"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	"github.com/pkg/errors"
)

const configfile = ".slack-incoming-webhooks-url"

func GetWebhookURL (path string) (uri string, err error) {
	if path == "" {
		uri, err = defaultConfig()
	} else if strings.Index(path, "http://") == 0 || strings.Index(path, "https://") == 0 {
		if _, err = url.ParseRequestURI(path); err == nil {
			uri = path
		}
	} else {
		if _, err = os.Stat(path); err == nil {
			uri, err = readConfig(path)
		}
	}

	return
}

func readConfig (path string) (url string, err error) {
	fp, err := os.Open(path)
	if err != nil {
		return "", errors.Wrap(err, "could not open config file")
	}
	defer fp.Close()

	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		url = scanner.Text()
		if url != "" {
			break
		}
	}

	return url, scanner.Err()
}

func defaultConfig() (string, error) {
	if path, err := getDefaultConfigFilePath(); err == nil {
		return readConfig(path)
	} else {
		return "", err
	}
}

func getDefaultConfigFilePath() (string, error) {
	home := os.Getenv("HOME")

	if home == "" {
		return "", errors.New("$HOME not set")
	}

	return filepath.Join(home, configfile), nil
}
