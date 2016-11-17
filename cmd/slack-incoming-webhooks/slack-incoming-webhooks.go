package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/hatyuki/go-slack-incoming-webhooks"
	"github.com/jessevdk/go-flags"
	"github.com/mattn/go-isatty"
)

const (
	exitOK = iota
	exitErr
	exitArgsParseErr = 255
)

var build string = "devel"

type options struct {
	WebHookURL string `long:"webhook-url" description:"Webhook URL to use."`
	Channel    string `short:"c" long:"channel" description:"channel the message should be sent to."`
	Username   string `short:"u" long:"username" description:"username that should be used as the sender."`
	IconEmoji  string `short:"i" long:"icon-emoji" description:"Slack emoji to use as the icon, e.g. ':ghost:'."`
	IconUrl    string `long:"icon-url" description:"URL of an icon image to use."`
	Help       bool   `short:"h" long:"help" description:"show this help message."`
}

func main() {
	os.Exit(Run())
}

func Run() int {
	opts, remain, err := parseOptions(os.Args[1:])
	if err != nil {
		return exitArgsParseErr
	} else if opts.Help {
		return exitOK
	}

	text, err := readText(remain[0:])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return exitArgsParseErr
	}

	client, err := slackIncomingWebhooks.NewClient(opts.WebHookURL)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return exitErr
	}

	payload := &slackIncomingWebhooks.Payload{
		Channel:   opts.Channel,
		Username:  opts.Username,
		IconEmoji: opts.IconEmoji,
		IconURL:   opts.IconUrl,
		Text:      text,
	}

	if err := client.Post(payload); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return exitErr
	}

	return exitOK
}

func parseOptions(args []string) (opts *options, remain []string, err error) {
	opts = &options{}
	parser := flags.NewParser(opts, flags.PrintErrors|flags.PassDoubleDash)
	parser.Usage = fmt.Sprintf("[OPTIONS] [MESSAGE ...]\n\nVersion:\n  %s (build: %s)", slackIncomingWebhooks.Version, build)

	if remain, err = parser.ParseArgs(args); opts.Help || err != nil {
		parser.WriteHelp(os.Stderr)
	}

	return
}

func readText(args []string) (string, error) {
	if isatty.IsTerminal(os.Stdin.Fd()) {
		return strings.Join(args, " "), nil
	} else if in, err := ioutil.ReadAll(os.Stdin); err == nil {
		return string(in), nil
	} else {
		return "", err
	}
}
