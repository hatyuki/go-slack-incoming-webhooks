package slackIncomingWebhooks

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/goulash/xdg"
)

func TestDefaultConfig(t *testing.T) {
	if _, err := ReadConfig(""); err == nil {
		t.Error("should return an error")
	}

	dir, err := ioutil.TempDir("", "")
	if err != nil {
		t.Skipf("could not create directory: %v", err)
	}
	defer os.RemoveAll(dir)

	orig := xdg.Getenv
	defer func() { xdg.Getenv = orig }()

	xdg.Getenv = func(key string) string {
		if key == "XDG_CONFIG_HOME" {
			return dir
		} else if key == "XDG_CONFIG_DIRS" {
			return ""
		} else {
			return orig(key)
		}
	}

	config := filepath.Join(dir, configfile)
	if err := os.MkdirAll(filepath.Dir(config), 0700); err != nil {
		t.Skipf("could not create directory: %v", err)
	}
}

func TestWebhookURL(t *testing.T) {
	want := "https://example.com/"

	if got, err := ReadConfig(want); err != nil {
		t.Errorf("returns an error: got %v", err)
	} else if got != want {
		t.Errorf("returns wrong URL: got %v want %v\n", got, want)
	}
}

func TestReadConfig(t *testing.T) {
	fp, err := ioutil.TempFile("", "")
	if err != nil {
		t.Skipf("could not create file: %v", err)
	}

	cleanup := func() {
		fp.Close()
		os.Remove(fp.Name())
	}
	defer cleanup()

	if _, err := ReadConfig(fp.Name()); err == nil {
		t.Error("should return an error")
	}

	want := "https://example.com/"
	fp.Write([]byte(want))
	if got, err := ReadConfig(fp.Name()); err != nil {
		t.Errorf("returns an error: got %v", err)
	} else if got != want {
		t.Errorf("returns wrong URL: got %v want %v\n", got, want)
	}

	cleanup()
	if _, err := ReadConfig(fp.Name()); err == nil {
		t.Error("should return an error")
	}
}
