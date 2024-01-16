package sendmsg

import "context"

// Sender is the interface for sending messages to specific message target like email, sms, etc.
type Sender interface {
	Send(ctx context.Context, message Message) error
}
