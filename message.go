package sendmsg

import "context"

type Message interface {
	GetRecipients() []string
	GetCC() []string
	GetBCC() []string

	GetSubject(ctx context.Context, vars map[string]any) (string, error)
	GetHTML(ctx context.Context, vars map[string]any) (string, error)
	GetPlainText(ctx context.Context, vars map[string]any) (string, error)
	GetAttaches() []Attach

	Complete() error
}
