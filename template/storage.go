package template

import (
	"context"
	"errors"
)

// ErrTemplateNotFound is the error returned when the template not found
var ErrTemplateNotFound = errors.New("template not found")

// Storage is the storage for the templates and access to them by name
type Storage interface {
	Template(ctx context.Context, name string) (Template, error)
}
