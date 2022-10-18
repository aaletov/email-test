package mailmgr

import (
	"net/smtp"

	"github.com/aaletov/email-test/pkg/topic"
)

const (
	Host       = "smtp.mailgun.org"
	Port       = "587"
	SMTPServer = Host + ":" + Port
)

type MailManager struct {
	User     string
	Password string
	Topics   []topic.Topic
}

func (m *MailManager) Start() {
	m.send(m.Topics[0])
}

func (m MailManager) send(topic topic.Topic) {
	for _, sub := range topic.Subscribers {
		to := []string{sub.Email}
		subject := topic.Subject
		body := "This is the body of the mail"
		message := []byte(subject + body)

		auth := smtp.PlainAuth("", m.User, m.Password, Host)

		err := smtp.SendMail(SMTPServer, auth, m.User, to, message)
		if err != nil {
			panic(err)
		}
	}
}
