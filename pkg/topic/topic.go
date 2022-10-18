package topic

import (
	sub "github.com/aaletov/email-test/pkg/subscriber"
)

type Topic struct {
	Name        string
	Subject     string
	Subscribers []sub.Subscriber
}
