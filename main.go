package main

import (
	"flag"

	"github.com/aaletov/email-test/pkg/mailmgr"
	"github.com/aaletov/email-test/pkg/subscriber"
	"github.com/aaletov/email-test/pkg/topic"
)

var (
	user     = flag.String("user", "", "User for SMTP server auth")
	password = flag.String("password", "", "Password for SMTP server auth")
)

func main() {
	flag.Parse()

	sub := subscriber.Subscriber{
		Email: "yabadabadushka@gmail.com",
	}
	myTopic := topic.Topic{
		Name:        "Cooking",
		Subject:     "Pan",
		Subscribers: []subscriber.Subscriber{sub},
	}
	mgr := mailmgr.MailManager{
		User:     *user,
		Password: *password,
		Topics:   []topic.Topic{myTopic},
	}
	mgr.Start()
}
