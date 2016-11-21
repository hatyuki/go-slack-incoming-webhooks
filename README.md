# slack-incoming-webhooks
Post messages to [Slack](https://slack.com) via [Incoming Webhooks](https://api.slack.com/incoming-webhooks).

## Installation
Using `go get`:

```bash
go get github.com/hatyuki/go-slack-incoming-webhooks
```

## Prerequisites
You need to create and configure an incoming webhook [here](https://my.slack.com/services/new/incoming-webhook/).
Grab the URL from the Webhook URL field.

### Configure Incoming Webhook URL
First, run following command:

```bash
slack-incoming-webhooks --configure "<YOUR_WEBHOOK_URL>"
```

or please run the command with `--webhook-url` option.

## Usage

    slack-incoming-webhooks [--webhook-url=WEBHOOK_URL] [-c=CHANNEL] [-u=USERNAME] [-e=ICON_EMOJI|-i=ICON_URL] [MESSAGE ...]
                            [--configure=WEBHOOK_URL]
                            [--help]

### Options
|Option           |Description                                   |
|-----------------|----------------------------------------------|
|-w, --webhook-url|Webhook URL to use                            |
|-c, --channel    |channel the message should be sent to         |
|-u, --username   |username that should be used as the sender    |
|-e, --icon-emoji |Slack emoji to use as the icon, e.g. `:ghost:`|
|-i, --icon-url   |URL of an icon image to use                   |
|    --configure  |set Webhook URL as default                    |
|-h, --help       |show help message                             |

### Examples
```bash
slack-incoming-webhooks --channel general "Hello, world!"
slack-incoming-webhooks --webhook-url "<YOUR_WEBHOOK_URL>" --username "I am bot" "Hello, incoming webhooks"
```

You can also pass in the message through _stdin_ like this:

```bash
echo -n "Hello, world!" | slack-incoming-webhooks --username bot --icon-emoji :ghost:
```
