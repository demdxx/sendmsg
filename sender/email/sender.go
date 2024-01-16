package email

import (
	"context"

	"github.com/ainsleyclark/go-mail/mail"

	"gitgub.com/demdxx/sendmsg"
)

// Sender implements sendmsg.Sender interface for email
type Sender struct {
	mailer mail.Mailer

	driver         string
	defaultVars    map[string]any
	defaultHeaders map[string]string
}

// New creates new email sender
func New(opts ...Option) (*Sender, error) {
	var o Options
	for _, opt := range opts {
		opt(&o)
	}
	mailer, err := o.Mailer()
	if err != nil {
		return nil, err
	}
	return &Sender{
		driver:         o.driver,
		mailer:         mailer,
		defaultVars:    o.vars,
		defaultHeaders: o.headers,
	}, nil
}

// Send sends message to the target
func (s *Sender) Send(ctx context.Context, message sendmsg.Message) error {
	tx, err := s.transmitionFromMessage(ctx, message)
	if err == nil {
		_, err = s.mailer.Send(tx)
	}
	return err
}

func (s *Sender) transmitionFromMessage(ctx context.Context, message sendmsg.Message) (*mail.Transmission, error) {
	subject, err := message.GetSubject(ctx, s.defaultVars)
	if err != nil {
		return nil, err
	}
	htmlText, err := message.GetHTML(ctx, s.defaultVars)
	if err != nil {
		return nil, err
	}
	plainText, err := message.GetPlainText(ctx, s.defaultVars)
	if err != nil {
		return nil, err
	}
	tx := &mail.Transmission{
		Recipients: message.GetRecipients(),
		CC:         message.GetCC(),
		BCC:        message.GetBCC(),
		Subject:    subject,
		HTML:       htmlText,
		PlainText:  plainText,
		Headers:    s.defaultHeaders,
	}
	for _, attach := range message.GetAttaches() {
		tx.Attachments = append(tx.Attachments, mail.Attachment{
			Filename: attach.Name(),
			Bytes:    attach.Content(),
		})
	}
	return tx, nil
}
