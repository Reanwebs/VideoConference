package utility

import (
	"fmt"
	"net/smtp"
	"time"

	"github.com/jordan-wright/email"
)

const (
	smtpAuthAddress   = "smtp.gmail.com"
	smtpServerAddress = "smtp.gmail.com:587"
)

type EmailSender interface {
	SendEmail(*ScheduleEmail) error
	MakeHostContent(conferenceType string, Time time.Time, Title string, Agenda string, Id string, Regards string) (string, error)
	MakeConferenceContent(conferenceType string, Time time.Time, Title string, Agenda string, Id string, Regards string) (string, error)
}

type GmailSender struct {
	name              string
	fromEmailAddress  string
	fromEmailPassword string
}

func NewGmailSender(name string, fromEmailAddress string, fromEmailPassword string) EmailSender {
	return &GmailSender{
		name:              name,
		fromEmailAddress:  fromEmailAddress,
		fromEmailPassword: fromEmailPassword,
	}
}

func (sender GmailSender) SendEmail(input *ScheduleEmail) error {
	e := email.NewEmail()
	e.From = fmt.Sprintf("%s <%s>", sender.name, sender.fromEmailAddress)
	e.Subject = input.Subject
	e.HTML = []byte(input.Content)
	e.To = input.To
	e.Cc = input.Cc
	e.Bcc = input.Bcc

	for _, f := range input.AttachFiles {
		_, err := e.AttachFile(f)
		if err != nil {
			return fmt.Errorf("failed to attach file %s: %w", f, err)
		}
	}

	smtpAuth := smtp.PlainAuth("", sender.fromEmailAddress, sender.fromEmailPassword, smtpAuthAddress)
	return e.Send(smtpServerAddress, smtpAuth)
}

func (sender GmailSender) MakeHostContent(conferenceType string, Time time.Time, Title string, Agenda string, id string, Regards string) (string, error) {
	TimeStr := Time.Format("2006-01-02 15:04:05")

	Content := `
Hi,

We are excited to confirm that your conference has been successfully scheduled on our website.
Details:
- ` + conferenceType + `
- ` + TimeStr + `
- ` + Agenda + `
- ` + id + `

If you have any questions or need to make any changes, please feel free to reach out to us. We're here to assist you every step of the way.

Thank you for choosing our platform to schedule your conference. We look forward to a successful event!

Best regards,
` + Regards + `
`

	return Content, nil
}

func (sender GmailSender) MakeConferenceContent(conferenceType string, Time time.Time, Title string, Agenda string, id string, Regards string) (string, error) {
	TimeStr := Time.Format("2006-01-02 15:04:05")

	Content := `
Hi,

We are excited to confirm that your conference has been successfully started on our website.
Details:
- ` + conferenceType + `
- ` + TimeStr + `
- ` + Agenda + `
- ` + id + `

You can invite participants through this link http://localhost:3000/media-container/` + id + `

If you have any questions or need to make any changes, please feel free to reach out to us. We're here to assist you every step of the way.

Thank you for choosing our platform for your conference. We look forward to a successful event!

Best regards,
` + Regards + `
`

	return Content, nil
}
