package slackIncomingWebhooks

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func TestConfigFromFile(t *testing.T) {
	fp, err := ioutil.TempFile("", "")
	if err != nil {
		t.Skip(err.Error())
	}
	defer func() {
		fp.Close()
		os.Remove(fp.Name())
	}()

	want := "http://example.com/"
	fp.Write([]byte(want))

	got, err := GetWebhookURL(fp.Name())
	if err != nil {
		t.Error(err.Error())
	} else if got != want {
		t.Errorf("returns wrong URL: got %v want %v\n", got, want)
	}
}

func TestNoConfigFile(t *testing.T) {
	if _, err := GetWebhookURL("/does/not/exists"); err == nil {
		t.Error("does not return error")
	}
}

func TestConfigURL(t *testing.T) {
	wants := [...]string{"http://example.com/", "https://example.com/"}

	for _, want := range wants {
		got, err := GetWebhookURL(want)
		if err != err {
			t.Errorf("returns error: got %v", err)
		} else if got != want {
			t.Errorf("returns wrong URL: got %v want %v\n", got, want)
		}
	}
}

func TestConfigInvalidURL(t *testing.T) {
	urls := [...]string{"ftp://example.com/", "http/https"}

	for _, url := range urls {
		if _, err := GetWebhookURL(url); err == nil {
			t.Error("does not return error")
		}
	}
}

func TestDefaultConfig(t *testing.T) {
	dir, err := ioutil.TempDir("", "")
	if err != nil {
		t.Skip(err.Error())
	}
	defer os.RemoveAll(dir)

	fp, err := os.Create(filepath.Join(dir, configfile))
	if err != nil {
		t.Skip(err.Error())
	}

	home := os.Getenv("HOME")
	defer os.Setenv("HOME", home)
	os.Setenv("HOME", dir)

	want := "http://example.com/"
	fp.Write([]byte(want))

	got, err := GetWebhookURL("")
	if err != nil {
		t.Errorf("returns error: got %v", err.Error())
	} else if got != want {
		t.Errorf("returns wrong URL: got %v want %v", got, want)
	}
}

func TestNoHome(t *testing.T) {
	home := os.Getenv("HOME")
	defer os.Setenv("HOME", home)
	os.Unsetenv("HOME")

	if _, err := GetWebhookURL(""); err == nil {
		t.Error("does not return error")
	}
}
