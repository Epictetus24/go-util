package notify

/*
POST an HTTPS request to https://api.pushover.net/1/messages.json with the following parameters:

	token - your application's API token (required)
	user - your user/group key (or that of your target user), viewable when logged into our dashboard; often referred to as USER_KEY in our documentation and code examples (required)
	message - your message (required)

Some optional parameters may also be included:

	attachment - an image attachment to send with the message (documentation)
	device - the name of one of your devices to send just to that device instead of all devices (documentation)
	html - set to 1 to enable HTML parsing (documentation)
	priority - a value of -2, -1, 0 (default), 1, or 2 (documentation)
	sound - the name of a supported sound to override your default sound choice (documentation)
	timestamp - a Unix timestamp of a time to display instead of when our API received it (documentation)
	title - your message's title, otherwise your app's name is used
	url - a supplementary URL to show with your message (documentation)
	url_title - a title for the URL specified as the url parameter, otherwise just the URL is shown (documentation)
*/
type POConf struct {
	Token string
	Users []string
}

type POmsg struct {
	Notifymsg
	Config     POConf
	Priority   int
	Device     string
	Html       int
	Attachment []byte
	Timestamp  string
	Sound      string
	Title      string
	Url        string
	Urltitle   string
}

func (p *POmsg) Send() {

}
