package dummy

import (
	"context"

	"gitgub.com/demdxx/sendmsg"
)

type Sender struct {
}

func NewSender() *Sender {
	return &Sender{}
}

func (s *Sender) Send(ctx context.Context, message sendmsg.Message) error {
	return nil
}
