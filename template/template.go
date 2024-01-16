package template

import "context"

// Template of the message to send
type Template interface {
	Name() string
	IsHTML() bool
	Subject(ctx context.Context, vars map[string]any) (string, error)
	Render(ctx context.Context, vars map[string]any) (string, error)
}
