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
	notExit = iota - 1
	exitOK
	exitErr
	exitArgsParseErr = 255
)

var build string = "devel"

type options struct {
	WebHookURL string `short:"w" long:"webhook-url" description:"Webhook URL to use."`
	Channel    string `short:"c" long:"channel" description:"channel the message should be sent to."`
	Username   string `short:"u" long:"username" description:"username that should be used as the sender."`
	IconEmoji  string `short:"e" long:"icon-emoji" description:"Slack emoji to use as the icon, e.g. ':ghost:'."`
	IconUrl    string `short:"i" long:"icon-url" description:"URL of an icon image to use."`
	Configure  string `long:"configure" description:"set Webhook URL as default."`
	Help       bool   `short:"h" long:"help" description:"show this help message."`
}

func main() {
	os.Exit(Run())
}

func Run() int {
	text, opts, status := parseArgs(os.Args[1:])
	if status != notExit {
		return status
	}

	url, err := slackIncomingWebhooks.ReadConfig(opts.WebHookURL)
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

	if err := slackIncomingWebhooks.Post(url, payload); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return exitErr
	}

	return exitOK
}

func parseArgs (args []string) (string, *options, int) {
	opts := &options{}
	parser := flags.NewParser(opts, flags.PrintErrors|flags.PassDoubleDash)
	parser.Usage = fmt.Sprintf("[OPTIONS] [MESSAGE ...]\n\nVersion:\n  %s (build: %s)", slackIncomingWebhooks.Version, build)

	remains, err := parser.ParseArgs(args)
	if err != nil {
		parser.WriteHelp(os.Stderr)
		return "", nil, exitArgsParseErr
	}

	if opts.Help {
		parser.WriteHelp(os.Stderr)
		return "", nil, exitOK
	}

	if opts.Configure != "" {
		if err := slackIncomingWebhooks.WriteConfig(opts.Configure); err != nil {
			fmt.Fprintln(os.Stderr, err)
			return "", nil, exitErr
		} else {
			return "", nil, exitOK
		}
	}

	if isatty.IsTerminal(os.Stdin.Fd()) {
		if len(remains) == 0 {
			parser.WriteHelp(os.Stderr)
			return "", nil, exitOK
		} else {
			return strings.Join(remains, " "), opts, notExit
		}
	} else {
		if in, err := ioutil.ReadAll(os.Stdin); err != nil {
			fmt.Fprintln(os.Stderr, err)
			return "", nil, exitArgsParseErr
		} else if text := string(in); text == "" {
			fmt.Fprintln(os.Stderr, "message body should not be empty")
			return "", nil, exitArgsParseErr
		} else {
			return text, opts, notExit
		}
	}
}
