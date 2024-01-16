package wrapper

import (
	"context"

	"gitgub.com/demdxx/sendmsg"
)

type Sender func(ctx context.Context, message sendmsg.Message) error

func (s Sender) Send(ctx context.Context, message sendmsg.Message) error {
	return s(ctx, message)
}
