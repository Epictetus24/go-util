package notify

/*
Simple functions for pushing to pushover, will likely add more for webhooks down the line.

*/

type Sender interface {
	Send(Notifymsg) error
}

type Notifymsg struct {
	MSGType string
	Msg     string
}
