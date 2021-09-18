package libraries

import (
	"fmt"
	"strings"

	"github.com/sirupsen/logrus"
	"gopkg.in/gomail.v2"
)

// allow less security apps from gmail

type MailLibrary struct {
	MailHost     string
	MailPort     int
	MailUser     string
	MailPassword string
}

func (lib MailLibrary) SendMail(subject, from, body string, to []string) {
	m := gomail.NewMessage()
	m.SetHeader("From", fmt.Sprintf("%s <%s>", from, lib.MailUser))
	m.SetHeader("To", to...)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)

	d := gomail.NewDialer(lib.MailHost, lib.MailPort, lib.MailUser, lib.MailPassword)
	if err := d.DialAndSend(m); err != nil {
		logrus.WithFields(logrus.Fields{
			"use_case":      "sendMail",
			"specification": fmt.Sprintf("error sending mail to:[%s]", strings.Join(to, ",")),
		}).Error(err.Error())
	}

}
