package template

import (
	"context"
	"io"
)

// Template of the message to send
type Template interface {
	Name() string
	IsHTML() bool
	Subject(ctx context.Context, vars map[string]any) (string, error)
	Render(ctx context.Context, wr io.Writer, vars map[string]any) error
}
