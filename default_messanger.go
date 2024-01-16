package sendmsg

import (
	"context"

	"gitgub.com/demdxx/sendmsg/template"
)

// DefaultMessanger implements sendmsg.Messanger interface
type DefaultMessanger struct {
	templateStore template.Storage
	senders       map[string]Sender
}

// NewDefaultMessanger creates new messanger
func NewDefaultMessanger(tmplStorage template.Storage) *DefaultMessanger {
	return &DefaultMessanger{
		templateStore: tmplStorage,
		senders:       make(map[string]Sender),
	}
}

// RegisterSender registers new sender
func (m *DefaultMessanger) RegisterSender(name string, sender Sender) *DefaultMessanger {
	m.senders[name] = sender
	return m
}

// Send sends the message to all senders that are allowed
func (m *DefaultMessanger) Send(ctx context.Context, opts ...SendOption) error {
	if len(m.senders) == 0 {
		return ErrNoSenders
	}
	options := &SendOptions{}
	for _, opt := range opts {
		opt(options)
	}
	message, err := options.GetMessage(m.templateStore)
	if err != nil {
		return err
	}
	sent := false
	for name, sender := range m.senders {
		if options.IsAllowed(name) {
			if err := sender.Send(ctx, message); err != nil {
				return err
			}
			sent = true
		}
	}
	if !sent {
		return ErrMessageNotSent
	}
	return nil
}
