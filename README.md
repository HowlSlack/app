<img src="./com/logo.png" alt="logo" width="25%" align="right"  />

<h1>Howl Slack</h1>

<br /><br />
<br /><br />
<br /><br />
An awesome program that reads out loud messages from slack channels.
<br /><br />
<br /><br />
<br /><br />

# Setup

## AWS

Download the aws-cli for your system.
Then configure it

```shell
aws configure
```

The app use the Polly service from S3.

More info on pricing:

```
https://aws.amazon.com/polly/pricing/
```

But it's **cheap**.

## Slack

Go to

```
https://YOUR_SUBDOMAIN.slack.com/apps/manage/custom-integrations
```

and create a token for a bot.

Assuming you called your bot `howl`, create a channel or add it to an existing one.

<img src="./com/channel-creation.png" alt="channel-creation" width="50%" align="center" />


## build

Add the auth token in `config.json`

```json
{
  "SlackToken": "your token"
}
```

Build the project:

Go version needed: `go1.11.2`

```shell
make
```

Then just run the binary

```shell
./howl-slack
```

The program will output the messages in real-time from the slack channel.

# Credits (icon)

<div>Icons made by <a href="https://www.freepik.com/" title="Freepik">Freepik</a> from <a href="https://www.flaticon.com/" 			    title="Flaticon">www.flaticon.com</a> is licensed by <a href="http://creativecommons.org/licenses/by/3.0/" 			    title="Creative Commons BY 3.0" target="_blank">CC 3.0 BY</a></div>



# License

`Howl Slack` is licensed under the GPLv3 License, sponsored and supported by <a href="https://hypermoon.io/" rel="noopener" target="_blank">HyperMoon</a>.
