# slack-incoming-webhooks
Post messages to [Slack](https://slack.com) via [Incoming Webhooks](https://api.slack.com/incoming-webhooks).

## Usage

```bash
slack-incoming-webhooks --webhook-url "<YOUR_INCOMING_WEBHOOK_URL>" --channel general "Hello, world!"
```

If `--webhook-url` not given, reads Webhook URL from `~/.slack-incoming-webhook-url`.

You can also pass in the message through _stdin_ like this:

```bash
echo -n "Hello, world!" | slack-incoming-webhooks --username bot --icon-emoji :ghost:
```

### Options

|Option          |Description                                  |
|----------------|---------------------------------------------|
|--webhook-url   |Webhook URL to use                           |
|-c, --channel   |channel the message should be sent to        |
|-u, --username  |username that should be used as the sender   |
|-i, --icon-emoji|Slack emoji to use as the icon, e.g. `:ghost`|
|--icon-url      |URL of an icon image to use                  |
|-h, --help      |show help message                            |

