package template

import (
	"context"
)

// DefaultStorage is the default implementation of the template storage
type DefaultStorage struct {
	templates map[string]Template
}

// NewDefaultStorage returns new default template storage
func NewDefaultStorage(tmpls ...Template) *DefaultStorage {
	mtmpls := map[string]Template{}
	for _, tmpl := range tmpls {
		mtmpls[tmpl.Name()] = tmpl
	}
	return &DefaultStorage{
		templates: mtmpls,
	}
}

// RegisterTmpl registers the templates
func (s *DefaultStorage) RegisterTmpl(tmpls ...Template) {
	for _, tmpl := range tmpls {
		s.templates[tmpl.Name()] = tmpl
	}
}

// Template returns the template by name
func (s *DefaultStorage) Template(ctx context.Context, name string) (Template, error) {
	if tmpl, ok := s.templates[name]; ok {
		return tmpl, nil
	}
	return nil, ErrTemplateNotFound
}
