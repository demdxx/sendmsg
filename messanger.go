package sendmsg

import (
	"context"
	"errors"

	"gitgub.com/demdxx/sendmsg/template"
)

var (
	ErrTemplateStorageNotDefined = errors.New("template storage is not defined")
	ErrRecipientsNotDefined      = errors.New("recipients are not defined")
	ErrSubjectNotDefined         = errors.New("subject is not defined")
	ErrContentNotDefined         = errors.New("content is not defined")
	ErrNoSenders                 = errors.New("no senders defined")
	ErrMessageNotSent            = errors.New("message not sent")
)

type SendOptions struct {
	AllowedSenders    []string
	NotAllowedSenders []string

	Message      Message
	TemplateName string
	Vars         map[string]any

	Recipients []string
	CC         []string
	BCC        []string

	Subject  string
	HTML     string
	Text     string
	Attaches []Attach
}

// IsAllowed checks if the sender is allowed to send the message
func (o *SendOptions) IsAllowed(sender string) bool {
	for _, s := range o.AllowedSenders {
		if s == sender {
			return true
		}
	}
	for _, s := range o.NotAllowedSenders {
		if s == sender {
			return false
		}
	}
	return len(o.AllowedSenders) == 0
}

// GetMessage returns the message to send
func (o *SendOptions) GetMessage(storage template.Storage) (Message, error) {
	if o.Message != nil {
		return o.Message, nil
	}
	if len(o.Recipients) == 0 {
		return nil, ErrRecipientsNotDefined
	}
	if o.TemplateName != "" {
		if storage == nil {
			return nil, ErrTemplateStorageNotDefined
		}
		template, err := storage.Template(context.Background(), o.TemplateName)
		if err != nil {
			return nil, err
		}
		return &TemplateMessage{
			Template:   template,
			Recipients: o.Recipients,
			CC:         o.CC,
			BCC:        o.BCC,
			Attaches:   o.Attaches,
		}, nil
	}
	if o.Subject == "" {
		return nil, ErrSubjectNotDefined
	}
	if o.HTML == "" || o.Text == "" {
		return nil, ErrContentNotDefined
	}
	return &DefaultMessage{
		Recipients: o.Recipients,
		CC:         o.CC,
		BCC:        o.BCC,
		Subject:    o.Subject,
		HTML:       o.HTML,
		PlainText:  o.Text,
		Attaches:   o.Attaches,
	}, nil
}

type SendOption func(*SendOptions)

// WithSender sets the list of allowed senders
func WithSender(senders ...string) SendOption {
	return func(o *SendOptions) {
		o.AllowedSenders = senders
	}
}

// WithoutSender sets the list of not allowed senders
func WithoutSender(senders ...string) SendOption {
	return func(o *SendOptions) {
		o.NotAllowedSenders = senders
	}
}

// WithMessage sets the message to send.
func WithMessage(message Message) SendOption {
	return func(o *SendOptions) {
		o.Message = message
	}
}

// WithTemplate sets the template name to send.
func WithTemplate(name string) SendOption {
	return func(o *SendOptions) {
		o.TemplateName = name
		if o.Message != nil {
			panic("sendmsg: WithTemplate and WithMessage are mutually exclusive")
		}
	}
}

// WithRecipients sets the list of recipients
func WithRecipients(recipients, cc, bcc []string) SendOption {
	return func(o *SendOptions) {
		o.Recipients = recipients
		o.CC = cc
		o.BCC = bcc
		if len(recipients) > 0 && o.Message != nil {
			panic("sendmsg: WithRecipients and WithMessage are mutually exclusive")
		}
	}
}

// WithSubject sets the subject of the message
func WithSubject(subject string) SendOption {
	return func(o *SendOptions) {
		o.Subject = subject
		if o.Message != nil {
			panic("sendmsg: WithSubject and WithMessage are mutually exclusive")
		}
	}
}

// WithContent sets the content of the message
func WithContent(html, text string) SendOption {
	return func(o *SendOptions) {
		o.HTML = html
		o.Text = text
		if o.Message != nil {
			panic("sendmsg: WithContent and WithMessage are mutually exclusive")
		}
	}
}

// WithAttaches sets the list of attachments
func WithAttaches(attaches ...Attach) SendOption {
	return func(o *SendOptions) {
		o.Attaches = attaches
		if o.Message != nil {
			panic("sendmsg: WithAttaches and WithMessage are mutually exclusive")
		}
	}
}

// WithVars sets the list of variables
func WithVars(vars map[string]any) SendOption {
	return func(o *SendOptions) {
		o.Vars = vars
	}
}

// Messanger interface for sending messages into abstract channel
type Messanger interface {
	Send(ctx context.Context, opts ...SendOption) error
}
